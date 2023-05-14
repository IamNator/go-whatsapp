package go_whatsapp_test

import (
	"context"
	"testing"

	whatsapp "github.com/IamNator/go-whatsapp/v3"
)

type testApiCaller struct {
	network map[string]whatsapp.APIResponse
}

func (m *testApiCaller) Post(url string, data []byte, headers map[string]string) (*whatsapp.APIResponse, int, error) {
	response := m.network[url+string(data)]
	return &response, 200, nil
}

// TODO: write better tests

// TestSendText tests the SendText function of the WhatsApp client.
func TestSendText(t *testing.T) {
	apiCaller := &testApiCaller{}

	testCases := []struct {
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
			to:       "987654321",
			text:     "how are you?",
			response: &whatsapp.APIResponse{},
		},
	}

	for _, tc := range testCases {
		whatsappClient := &whatsapp.Client{}
		whatsappClient.SetApiCaller(apiCaller)

		_, err := whatsappClient.SendText(context.TODO(), tc.to, tc.text)
		if err != nil {
			t.Errorf("Failed to send text message: %v", err)
		}
	}
}
