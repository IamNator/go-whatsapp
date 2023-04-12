package go_whatsapp

import (
	"encoding/json"
	"strings"

	"github.com/iamNator/go-whatsapp/template"
)

type (
	// RequestPayload is the payload to be sent to the WhatsApp Cloud API
	RequestPayload struct {
		MessagingProduct string             `json:"messaging_product"` // e.g whatsapp
		To               string             `json:"to"`                // e.g 2349045057268
		Type             PayloadType        `json:"type"`              // e.g text, template
		Text             *Text              `json:"text,omitempty"`    // if type is text
		Template         *template.Template `json:"template"`          // if type is template
	}

	PayloadType string

	Text struct {
		Body string `json:"body"`
	}
)

const (
	TypeText     PayloadType = "text"
	TypeTemplate PayloadType = "template"

	whatsApp = "whatsapp"
)

func (m *RequestPayload) Byte() ([]byte, error) {
	return json.Marshal(m)
}

func removeLeadPlusSign(s string) string {
	return strings.ReplaceAll(s, "+", "")
}

func NewPayload(templateName, to string, langCode template.LanguageCode) *RequestPayload {

	return &RequestPayload{
		MessagingProduct: whatsApp,
		To:               removeLeadPlusSign(to),
		Type:             TypeTemplate,
		Template:         template.New(templateName, langCode),
	}
}

func NewFromBytes(b []byte) (*RequestPayload, error) {
	var m RequestPayload
	if er := json.Unmarshal(b, &m); er != nil {
		return nil, er
	}

	return &m, nil
}
