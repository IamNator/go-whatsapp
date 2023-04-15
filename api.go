package go_whatsapp

import (
	"context"
	"os"

	"github.com/iamNator/go-whatsapp/errors"

	"github.com/iamNator/go-whatsapp/template"
)

type (
	Client struct {
		phoneNumberID string
		accessToken   string
		baseURL       string
		apiVersion    MetaAPIVersion
		apiCaller     IApiCaller
	}
)

// New
//
// e.g  _meta := New(
//
//	         "9484589000430090",
//				"44NSNANSF094545nLKJGSJFSKF78985395495NKSJNFDJNSKFNSNJFNSDNFSDNFJNSDKFNSDJFNJSDNFJSD",
//	         V14 )
func New(phoneNumberID, metaAppAccessToken string, apiVersion MetaAPIVersion) *Client {

	baseURL := "https://graph.facebook.com"

	if baseU := os.Getenv("META_BASE_URL"); baseU != "" {
		baseURL = baseU
	}

	// 40 requests per second

	return &Client{
		phoneNumberID: phoneNumberID,
		accessToken:   metaAppAccessToken,
		baseURL:       baseURL,
		apiVersion:    apiVersion,
		apiCaller:     &apiCaller{}, // default
	}
}

func (m *Client) SetBaseURL(url string) {
	m.baseURL = url
}

func (m *Client) SetApiVersion(apiVersion MetaAPIVersion) {
	m.apiVersion = apiVersion
}

func (m *Client) SetApiCaller(apiCaller IApiCaller) {
	m.apiCaller = apiCaller
}

type (

	//WhatsappOutputError ..
	APIError struct {
		Message      string       `json:"message"`
		Type         string       `json:"type"`
		Code         int          `json:"code"`
		ErrorData    APIErrorData `json:"error_data"`
		ErrorSubCode uint         `json:"error_subcode"`
		FBTraceID    string       `json:"fbtrace_id"`
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

func (e APIError) Error() string {
	return errors.Error(e.Code).Error()
}

// Send sends a message
func (m *Client) Send(ctx context.Context, msg RequestPayload) (*APIResponse, *APIError, error) {

	url := m.baseURL + "/" + m.apiVersion.String() + "/" + m.phoneNumberID + "/messages"
	headers := map[string]string{
		"Authorization": "Bearer " + m.accessToken,
	}

	//convert to json
	data, er := msg.Byte()
	if er != nil {
		return nil, nil, er
	}

	output, statusCode, er := m.apiCaller.Post(
		url,
		data,
		headers)
	if er != nil {
		return nil, nil, er
	}

	//check for error
	if errors.IsErrorCode(output.Error.Code, statusCode) {
		return nil, &output.Error, nil
	}

	//check for error
	if output.Error.ErrorSubCode != 0 {
		return nil, &output.Error, nil
	}

	return output, nil, nil
}

// SendText sends a text message
func (m *Client) SendText(ctx context.Context, to string, text string) (*APIResponse, *APIError, error) {

	msg := NewPayloadWithText(to, text)

	url := m.baseURL + "/" + m.apiVersion.String() + "/" + m.phoneNumberID + "/messages"
	headers := map[string]string{
		"Authorization": "Bearer " + m.accessToken,
	}

	//convert to json
	data, er := msg.Byte()
	if er != nil {
		return nil, nil, er
	}

	output, statusCode, er := m.apiCaller.Post(
		url,
		data,
		headers)
	if er != nil {
		return nil, nil, er
	}

	//check for error
	if errors.IsErrorCode(output.Error.Code, statusCode) {
		return nil, &output.Error, nil
	}

	//check for error
	if output.Error.ErrorSubCode != 0 {
		return nil, &output.Error, nil
	}

	return output, nil, nil
}

// SendTemplate sends a template message
//
// APIResponse, *APIError, error
//
// APIResponse: response from the server
//
// *APIError: error from the server
//
// error: error from the client
func (m *Client) SendTemplate(ctx context.Context, to string, tmpl template.Template) (*APIResponse, *APIError, error) {

	msg := NewPayloadWithTemplate(to, tmpl)

	url := m.baseURL + "/" + m.apiVersion.String() + "/" + m.phoneNumberID + "/messages"
	headers := map[string]string{
		"Authorization": "Bearer " + m.accessToken,
	}

	//convert to json
	data, er := msg.Byte()
	if er != nil {
		return nil, nil, er
	}

	output, statusCode, er := m.apiCaller.Post(
		url,
		data,
		headers)
	if er != nil {
		return nil, nil, er
	}

	//check for error
	if errors.IsErrorCode(output.Error.Code, statusCode) {
		return nil, &output.Error, nil
	}

	return output, nil, nil
}

// ------------------------------------------------  REST CALLS -------------------------------
