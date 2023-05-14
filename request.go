package go_whatsapp

import (
	"encoding/json"
	"strings"

	"github.com/IamNator/go-whatsapp/v3/template"
)

type (
	// APIRequest is the payload to be sent to the WhatsApp Cloud API
	APIRequest struct {
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

func (m APIRequest) Byte() ([]byte, error) {
	return json.Marshal(m)
}

func removeLeadingPlusSign(s string) string {
	return strings.ReplaceAll(s, "+", "")
}

// Deprecated
func NewAPIRequest(templateName, to string, langCode template.LanguageCode) APIRequest {

	return APIRequest{
		MessagingProduct: whatsApp,
		To:               removeLeadingPlusSign(to),
		Type:             TypeTemplate,
		Template:         template.New(templateName, langCode),
	}
}

func NewAPIRequestWithText(to, text string) APIRequest {
	return APIRequest{
		MessagingProduct: whatsApp,
		To:               removeLeadingPlusSign(to),
		Type:             TypeText,
		Text: &Text{
			Body: text,
		},
	}
}

func NewAPIRequestWithTemplate(to string, tmpl template.Template) APIRequest {
	return APIRequest{
		MessagingProduct: whatsApp,
		To:               removeLeadingPlusSign(to),
		Type:             TypeTemplate,
		Template:         &tmpl,
	}
}

func NewAPIRequestFromBytes(b []byte) (*APIRequest, error) {
	var m APIRequest
	if er := json.Unmarshal(b, &m); er != nil {
		return nil, er
	}

	return &m, nil
}
