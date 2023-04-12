package go_whatsapp

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/iamNator/go-whatsapp/errors"
)

type (
	META struct {
		phoneNumberID string
		accessToken   string
		baseURL       string
		apiVersion    MetaAPIVersion
	}
)

// New
//
// e.g  _meta := New(
//
//	         "9484589000430090",
//				"44NSNANSF094545nLKJGSJFSKF78985395495NKSJNFDJNSKFNSNJFNSDNFSDNFJNSDKFNSDJFNJSDNFJSD",
//	         V14 )
func New(phoneNumberID, metaAppAccessToken string, apiVersion MetaAPIVersion) *META {

	baseURL := "https://graph.facebook.com"

	if baseU := os.Getenv("META_BASE_URL"); baseU != "" {
		baseURL = baseU
	}

	// 40 requests per second

	return &META{
		phoneNumberID: phoneNumberID,
		accessToken:   metaAppAccessToken,
		baseURL:       baseURL,
		apiVersion:    apiVersion,
	}
}

func (m *META) SetBaseURL(url string) {
	m.baseURL = url
}

type (

	//WhatsappOutputError ..
	APIError struct {
		Message      string           `json:"message"`
		Type         string           `json:"type"`
		Code         errors.MetaError `json:"code"`
		ErrorData    APIErrorData     `json:"error_data"`
		ErrorSubCode uint             `json:"error_subcode"`
		FBTraceID    string           `json:"fbtrace_id"`
	}

	APIErrorData struct {
		Details          string `json:"details"`
		MessagingProduct string `json:"messaging_product"`
	}

	APIResponseContact struct {
		Input string `json:"input"`
		WaID  string `json:"wa_id"`
	}

	APIResponseMessage struct {
		ID string `json:"id"`
	}

	APIResponse struct {
		Error            APIError             `json:"error"`
		MessagingProduct string               `json:"messaging_product"`
		Contacts         []APIResponseContact `json:"contacts"`
		Messages         []APIResponseMessage `json:"messages"`
	}
)

func (m *META) Send(ctx context.Context, msg RequestPayload) (*APIResponse, *APIError, error) {

	url := m.baseURL + "/" + m.apiVersion.String() + "/" + m.phoneNumberID + "/messages"
	headers := map[string]string{
		"Authorization": "Bearer " + m.accessToken,
	}

	//convert to json
	data, er := msg.Byte()
	if er != nil {
		return nil, nil, er
	}

	output, er := Post[APIResponse](
		url,
		data,
		headers)
	if er != nil {
		return nil, nil, er
	}

	//check for error
	if output.Error.ErrorSubCode != 0 {
		return nil, &output.Error, nil
	}

	return output, nil, nil
}

// ------------------------------------------------  REST CALLS -------------------------------

func Post[Response any](url string, data []byte, headers map[string]string) (*Response, error) {

	request, er := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if er != nil {
		return nil, er
	}

	client := http.DefaultClient
	response, er := client.Do(request)
	if er != nil {
		return nil, er
	}

	responseData, er := ioutil.ReadAll(response.Body)
	if er != nil {
		return nil, er
	}

	var output Response
	if er := json.Unmarshal(responseData, &output); er != nil {
		return nil, er
	}

	return &output, nil
}
