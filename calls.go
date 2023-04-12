package go_whatsapp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// IApiCaller is the interface for the API caller
//
// It interfaces with external network resources
type IApiCaller interface {
	Post(url string, data []byte, headers map[string]string) (*APIResponse, int, error)
}

type apiCaller struct {
	// contains filtered or unexported fields
}

func (m *apiCaller) Post(url string, body []byte, headers map[string]string) (*APIResponse, int, error) {
	addBearerToken(headers, headers["Authorization"])      // add bearer token to headers
	var response APIResponse                               // create a response object
	statusCode, err := post(url, body, headers, &response) // make the post request
	if err != nil {
		return nil, statusCode, err // return the error
	}
	return &response, statusCode, err // return the response object and the error
}

// addBearerToken adds the bearer token to the headers
func addBearerToken(headers map[string]string, token string) {
	headers["Authorization"] = "Bearer " + token
}

// attachHeaders attaches the headers to the request
func attachHeaders(request *http.Request, headers map[string]string) {
	for key, value := range headers {
		request.Header.Set(key, value)
	}
}

// post makes a post request to the given url
func post(url string, body []byte, headers map[string]string, response interface{}) (int, error) {

	request, er := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if er != nil {
		return 0, er
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+headers["Authorization"])
	attachHeaders(request, headers)

	client := http.DefaultClient
	apiResponse, er := client.Do(request)
	if er != nil {
		return 0, er
	}

	responseData, er := io.ReadAll(apiResponse.Body)
	if er != nil {
		return 0, er
	}

	if er := json.Unmarshal(responseData, response); er != nil {
		return 0, er
	}

	return apiResponse.StatusCode, nil
}
