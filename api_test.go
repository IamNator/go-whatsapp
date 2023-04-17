package go_whatsapp_test

import (
	"context"
	"testing"

	whatsapp "github.com/iamNator/go-whatsapp/v2"
)

type testApiCaller struct {
	network map[string]whatsapp.APIResponse
}

func (m *testApiCaller) Post(url string, data []byte, headers map[string]string) (*whatsapp.APIResponse, int, error) {
	response := m.network[url+string(data)]
	return &response, 200, nil
}

// TODO: write better tests
func TestSendText(t *testing.T) {

	apiCaller := &testApiCaller{}

	tt := []struct {
		to       string
		text     string
		response *whatsapp.APIResponse
	}{
		{
			to:       "123456789",
			text:     "hello world",
			response: &whatsapp.APIResponse{},
		},
		{
			to:       "123456789",
			text:     "hello world",
			response: &whatsapp.APIResponse{},
		},
	}

	for _, tc := range tt {
		m := &whatsapp.Client{}
		m.SetApiCaller(apiCaller)

		_, _, err := m.SendText(context.TODO(), tc.to, tc.text)
		if err != nil {
			t.Errorf("error: %v", err)
		}
	}
}
