package go_whatsapp

import (
	"context"

	"github.com/IamNator/go-whatsapp/v3/errors"
	"github.com/IamNator/go-whatsapp/v3/template"
)

type (
	Client struct {
		phoneNumberID string
		accessToken   string
		baseURL       string
		apiVersion    APIVersion
		apiCaller     APICaller
		debug         bool
	}
)

type Opt func(*Client)

// WithBaseURL sets the base url
func WithBaseURL(url string) Opt {
	return func(m *Client) {
		m.baseURL = url
	}
}

// WithApiVersion sets the api version
func WithApiVersion(apiVersion APIVersion) Opt {
	return func(m *Client) {
		m.apiVersion = apiVersion
	}
}

// WithApiCaller sets the api caller
func WithApiCaller(apiCaller APICaller) Opt {
	return func(m *Client) {
		m.apiCaller = apiCaller
	}
}

// New returns a new instance of the Client
//
// e.g  client := New(
//
//	         "9414589060430990",
//			 "44NSNAUSF094545nLKIGSJFSKF78985395495NKSJNFDJNS0FNSNJFNSDNFSDNFJNSDKFKSDJFNJSDNFJSD",
//	         WithApiVersion(V16), WithBaseURL("https://graph.facebook.com") )
//
//	phoneNumberID: the phone number id
//	appAccessToken: the app access token
//	opts: the options
//
// by default, the api version is V16
// and the base url is https://graph.facebook.com
func New(phoneNumberID, appAccessToken string, opts ...Opt) *Client {

	const baseURL = "https://graph.facebook.com"

	m := &Client{
		phoneNumberID: phoneNumberID,
		accessToken:   appAccessToken,
		baseURL:       baseURL,
		apiVersion:    V15,
		apiCaller:     &apiCaller{}, // default
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
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
		ErrorSubCode int          `json:"error_subcode"`
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
		Error            *APIError            `json:"error"`
		MessagingProduct string               `json:"messaging_product"` // e.g whatsapp
		Contacts         []APIResponseContact `json:"contacts"`
		Messages         []APIResponseMessage `json:"messages"`
	}
)

func (e APIError) Error() string {
	return errors.Error(e.Code).Error()
}

// Send sends a message
func (m *Client) Send(ctx context.Context, msg APIRequest) (*APIResponse, error) {

	url := m.baseURL + "/" + m.apiVersion.String() + "/" + m.phoneNumberID + "/messages"
	headers := map[string]string{
		"Authorization": "Bearer " + m.accessToken, // set the authorization header
		"Content-Type":  "application/json",        // set the content type header
		"Accept":        "application/json",        // set the accept header
	}

	//convert to bytes
	data, er := msg.Byte()
	if er != nil {
		return nil, er
	}

	output, _, er := m.apiCaller.Post(
		url,
		data,
		headers)
	if er != nil {
		return nil, er
	}

	return output, nil
}

// SendText sends a text message
func (m *Client) SendText(ctx context.Context, to string, text string) (*APIResponse, error) {

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
func (m *Client) SendTemplate(ctx context.Context, to string, tmpl template.Template) (*APIResponse, error) {

	msg := NewAPIRequestWithTemplate(to, tmpl) // create an api request payload with template

	return m.Send(ctx, msg)
}
