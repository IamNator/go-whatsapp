package go_whatsapp

import (
	"context"

	"github.com/IamNator/go-whatsapp/v2/errors"

	"github.com/IamNator/go-whatsapp/v2/template"
)

type (
	Client struct {
		phoneNumberID string
		accessToken   string
		baseURL       string
		apiVersion    APIVersion
		apiCaller     APICaller
	}
)

// New
//
// e.g  _meta := New(
//
//	         "9484589000430090",
//				"44NSNANSF094545nLKJGSJFSKF78985395495NKSJNFDJNSKFNSNJFNSDNFSDNFJNSDKFNSDJFNJSDNFJSD",
//	         V14 )

func New(phoneNumberID, metaAppAccessToken string, apiVersion APIVersion) *Client {
	const baseURL = "https://graph.facebook.com"

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

func (m *Client) SetApiVersion(apiVersion APIVersion) {
	m.apiVersion = apiVersion
}

func (m *Client) SetApiCaller(apiCaller APICaller) {
	m.apiCaller = apiCaller
}

type (
	APIErrorData struct {
		Details          string `json:"details"`
		MessagingProduct string `json:"messaging_product"`
	}

	APIError struct {
		Message      string       `json:"message"`
		Type         string       `json:"type"`
		Code         int          `json:"code"`
		ErrorData    APIErrorData `json:"error_data"`
		ErrorSubCode uint         `json:"error_subcode"`
		FBTraceID    string       `json:"fbtrace_id"`
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
		MessagingProduct string               `json:"messaging_product"` // e.g whatsapp
		Contacts         []APIResponseContact `json:"contacts"` 
		Messages         []APIResponseMessage `json:"messages"`
	}
)

func (e APIError) Error() string {
	return errors.Error(e.Code).Error()
}

// Send sends a message
func (m *Client) Send(ctx context.Context, msg APIRequest) (*APIResponse, *APIError, error) {

	url := m.baseURL + "/" + m.apiVersion.String() + "/" + m.phoneNumberID + "/messages"
	headers := map[string]string{
		"Authorization": "Bearer " + m.accessToken,
	}

	//convert to bytes
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

	msg := NewAPIRequestWithText(to, text) // create an api request payload with text

	return m.Send(ctx, msg)
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

	msg := NewAPIRequestWithTemplate(to, tmpl) // create an api request payload with template

	return m.Send(ctx, msg)
}
