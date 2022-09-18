package go_whatsapp

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/iamNator/go-whatsapp/template"
	"io/ioutil"
	"net/http"
	"os"
)

type (
	META struct {
		phoneNumberID string
		accessToken   string
		baseURL       string
		apiVersion    MetaAPIVersion

		storagePlugin IStoragePlugin
	}
)

// New
//
// e.g  _meta := New(
//	         "9484589000430090",
//				"44NSNANSF094545nLKJGSJFSKF78985395495NKSJNFDJNSKFNSNJFNSDNFSDNFJNSDKFNSDJFNJSDNFJSD",
//	         V14 )
//
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
		storagePlugin: nil,
	}
}

func (m *META) AttachStoragePlugin(storagePlugin IStoragePlugin) {
	m.storagePlugin = storagePlugin
}

func (m *META) GetStoragePlugin() IStoragePlugin {
	return m.storagePlugin
}

func (m META) CheckStorageExist() bool {
	return m.storagePlugin != nil
}

func (m *META) SetBaseURL(url string) {
	m.baseURL = url
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

type (
	WhatsappOutput struct {
		Error            WhatsappOutputError     `json:"error"`
		MessagingProduct string                  `json:"messaging_product"`
		Contacts         []WhatsappOutputContact `json:"contacts"`
		Messages         []WhatsappOutputMessage `json:"messages"`
	}

	//WhatsappOutputError ..
	WhatsappOutputError struct {
		ErrorSubCode uint                    `json:"error_subcode"`
		ErrorData    WhatsappOutputErrorData `json:"error_data"`
	}

	WhatsappOutputErrorData struct {
		Details string `json:"details"`
	}

	WhatsappOutputContact struct {
		Input string `json:"input"`
		WaID  string `json:"wa_id"`
	}

	WhatsappOutputMessage struct {
		ID string `json:"id"`
	}
)

func (m *META) Send(ctx context.Context, msg Message) (*Response, *WhatsappOutputError, error) {

	payload, er := template.FromByteToMap(
		msg.Data,
	)
	if er != nil {
		return nil, nil, er
	}

	url := m.baseURL + "/" + m.apiVersion.String() + "/" + m.phoneNumberID + "/messages"
	headers := map[string]string{
		"Authorization": "Bearer " + m.accessToken,
	}

	output, er := post(url, payload, headers)
	if er != nil {
		return nil, nil, er
	}

	//check for error
	if output.Error.ErrorSubCode != 0 {
		return nil, &output.Error, nil
	}

	return &Response{
		Status: "success",
	}, nil, nil
}

// ------------------------------------------------  REST CALLS -------------------------------

func post(url string, data map[string]interface{}, headers map[string]string) (*WhatsappOutput, error) {

	rawData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(rawData)
	request, er := http.NewRequest(http.MethodPost, url, reader)
	if er != nil {
		return nil, er
	}

	client := http.DefaultClient
	response, er := client.Do(request)
	if er != nil {
		return nil, er
	}

	output := new(WhatsappOutput)
	responseData, er := ioutil.ReadAll(response.Body)
	if er != nil {
		return nil, er
	}

	if er := json.Unmarshal(responseData, output); er != nil {
		return nil, er
	}

	return output, nil
}

// --------------------------------------------------   STORAGE  -----------------------------------------------

func makeKey(key string) string {
	return "whatsapp_key_" + key + "_whatsapp_key"
}
func (m *META) save(key string, value interface{}) error {
	if !m.CheckStorageExist() {
		return nil
	}
	return m.storagePlugin.Store(makeKey(key), value)
}
func (m *META) get(key string) (interface{}, error) {
	if !m.CheckStorageExist() {
		return "", nil
	}
	return m.storagePlugin.Load(makeKey(key))
}
func (m *META) drop(key string) error {
	if !m.CheckStorageExist() {
		return nil
	}
	return m.storagePlugin.Remove(makeKey(key))
}
