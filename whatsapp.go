package go_whatsapp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	resty "github.com/go-resty/resty/v2"
	"github.com/iamNator/go-whatsapp/meta"
	"os"
	"time"

	"golang.org/x/time/rate"
)

type META struct {
	client      *resty.Client
	rateLimiter *rate.Limiter
	appId       string
	accessToken string
	baseURL     string
}

func New(metaAppId, metaAppAccessToken string) *META {

	baseURL := "https://graph.facebook.com"

	if baseU := os.Getenv("META_BASE_URL"); baseU != "" {
		baseURL = baseU
	}

	client := resty.New()
	client.EnableTrace()
	client.SetRetryCount(3)
	client.SetRetryWaitTime(300 * time.Millisecond)
	client.SetTimeout(time.Second)
	client.SetRetryMaxWaitTime(time.Second)
	client.SetBaseURL(baseURL)

	rateLimiter := rate.NewLimiter(rate.Every(time.Second), 40) // 40 requests per second

	return &META{
		rateLimiter: rateLimiter,
		client:      client,
		appId:       metaAppId,
		accessToken: metaAppAccessToken,
		baseURL:     baseURL,
	}
}

type (
	//nolint
	metaPayload struct {
		MessagingProduct string               `json:"messaging_product"`
		To               string               `json:"to"`
		Type             string               `json:"type"`
		Text             *Text                `json:"text"`
		Template         *MetaPayloadTemplate `json:"template"`
	}

	Text struct {
		Body string `json:"body"`
	}

	MetaPayloadTemplate struct {
		Name       string      `json:"name"`
		Language   *Language   `json:"language"`
		Components []Component `json:"components"`
	}

	Language struct {
		Code string `json:"code"`
	}

	Component struct {
		Type       string      `json:"type"`
		Parameters []Parameter `json:"parameters"`
	}

	Parameter struct {
		Type string `json:"type"`
		Text string `json:"text"`
	}

	WhatsAppErrType struct {
		MessagingProduct string `json:"messaging_product"`
		Contacts         []struct {
			Input string `json:"input"`
			WaId  string `json:"wa_id"`
		} `json:"contacts"`
		Messages []struct {
			Id string `json:"id"`
		} `json:"messages"`
	}

	Message struct {
		Data []byte `json:"data"`
	}

	Response struct {
		Status string `json:"status"`
	}
)

const (
	TypeText     string = "text"
	TypeTemplate string = "template"
)

func (m *META) Send(ctx context.Context, msg Message) (*Response, error) {

	payload, er := meta.FromByteToMap(
		msg.Data,
	)
	if er != nil {
		return nil, er
	}

	output := struct {
		Error struct {
			ErrorSubcode uint `json:"error_subcode"`
			ErrorData    struct {
				Details string `json:"details"`
			} `json:"error_data"`
		} `json:"error"`
		MessagingProduct string `json:"messaging_product"`
		Contacts         []struct {
			Input string `json:"input"`
			WaID  string `json:"wa_id"`
		} `json:"contacts"`
		Messages []struct {
			ID string `json:"id"`
		} `json:"messages"`
	}{}

	err := m.rateLimiter.Wait(ctx) // This is a blocking call. Honors the rate limit
	if err != nil {
		return nil, err
	}
	resp, err := m.client.
		R().
		SetBody(payload).
		EnableTrace().
		SetHeader("Authorization", "Bearer "+m.accessToken).
		SetAuthToken(m.accessToken).
		Post("/v13.0/" + m.appId + "/messages")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() > 300 {
		return nil, fmt.Errorf("%s", resp.String())
	}

	if er := json.Unmarshal([]byte(resp.String()), &output); er != nil {
		return nil, er
	}

	//chech for error
	if output.Error.ErrorData.Details != "" {
		return nil, errors.New(output.Error.ErrorData.Details)
	}

	return &Response{
		Status: "success",
	}, nil
}
