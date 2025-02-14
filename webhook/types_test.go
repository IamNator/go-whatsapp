package webhook

import (
	"os"
	"testing"
)

// TestParseWebhook tests webhook parsing with different message types
func TestParseWebhook(t *testing.T) {
	// Test cases with their respective JSON files
	tests := []struct {
		name     string
		jsonPath string
		wantType string
	}{
		{
			name:     "Text Message",
			jsonPath: "testdata/text_message.json",
			wantType: "text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			jsonFile, err := os.ReadFile(tt.jsonPath)
			if err != nil {
				t.Fatalf("failed to read test file: %v", err)
			}

			webhook, err := ParseBytes(jsonFile)
			if err != nil {
				t.Fatalf("failed to parse webhook: %v", err)
			}

			// Basic validation
			if webhook.Object != "whatsapp_business_account" {
				t.Errorf("expected object to be whatsapp_business_account, got %s", webhook.Object)
			}

			// Validate message type
			if len(webhook.Entry) > 0 &&
				len(webhook.Entry[0].Changes) > 0 &&
				len(webhook.Entry[0].Changes[0].Value.Messages) > 0 {

				msgType := webhook.Entry[0].Changes[0].Value.Messages[0].Type
				if msgType != tt.wantType {
					t.Errorf("expected message type %s, got %s", tt.wantType, msgType)
				}

			} else {
				t.Error("webhook message structure is invalid")
			}
		})
	}
}
