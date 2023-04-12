package go_whatsapp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// IApiCaller is the interface for the API caller
//
// It interfaces with external network resources
type IApiCaller interface {
	Post(url string, data []byte, headers map[string]string) (*APIResponse, error)
}

type apiCaller struct {
	// contains filtered or unexported fields
}

func (m *apiCaller) Post(url string, data []byte, headers map[string]string) (*APIResponse, error) {
	return post(url, data, headers)
}

func post(url string, data []byte, headers map[string]string) (*APIResponse, error) {

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

	var output APIResponse
	if er := json.Unmarshal(responseData, &output); er != nil {
		return nil, er
	}

	return &output, nil
}
