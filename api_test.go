package go_whatsapp

import (
	"context"
	"testing"
)

type testApiCaller struct {
	network map[string]APIResponse
}

func (m *testApiCaller) Post(url string, data []byte, headers map[string]string) (*APIResponse, int, error) {
	response := m.network[url+string(data)]
	return &response, 200, nil
}

// TODO: write better tests
func TestSendText(t *testing.T) {

	apiCaller := &testApiCaller{}

	tt := []struct {
		to       string
		text     string
		response *APIResponse
	}{
		{
			to:       "123456789",
			text:     "hello world",
			response: &APIResponse{},
		},
		{
			to:       "123456789",
			text:     "hello world",
			response: &APIResponse{},
		},
	}

	for _, tc := range tt {
		m := &Client{
			apiCaller: apiCaller,
		}

		_, _, err := m.SendText(context.TODO(), tc.to, tc.text)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	}
}
