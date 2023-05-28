package go_whatsapp

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// IApiCaller is the interface for the API caller
//
// It interfaces with external network resources
type APICaller interface {
	Post(url string, data []byte, headers map[string]string) (response *APIResponse, statusCode int, err error)
}

type apiCaller struct {
	// contains filtered or unexported fields
}

func (m *apiCaller) Post(url string, body []byte, headers map[string]string) (*APIResponse, int, error) {
	var response APIResponse                               // create a response object
	statusCode, err := post(url, body, headers, &response) // make the post request
	if err != nil {
		return nil, statusCode, err // return the error
	}
	return &response, statusCode, nil // return the response object and the error
}

// attachHeaders attaches the headers to the request
func attachHeaders(request *http.Request, headers map[string]string) {
	for key, value := range headers {
		request.Header.Set(key, value)
	}
}

// post makes a post request to the given url
func post(url string, body []byte, headers map[string]string, response interface{}) (int, error) {

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return http.StatusBadRequest, err
	}

	attachHeaders(request, headers)

	client := http.DefaultClient
	apiResponse, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	defer apiResponse.Body.Close()

	if err := json.NewDecoder(apiResponse.Body).Decode(response); err != nil {
		return 0, err
	}

	return apiResponse.StatusCode, nil
}
