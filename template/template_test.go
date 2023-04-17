package template_test

import (
	"strings"
	"testing"

	"github.com/IamNator/go-whatsapp/v2/template"
)

type (
	components struct {
		Header []string
		Body   []string
		Button []string
	}
)

func TestBytes(t *testing.T) {

	tt := []struct {
		TemplateName string
		To           string
		Language     template.LanguageCode
		Components   components
		Output       string
	}{
		{
			TemplateName: "signup_otp_1",
			Language:     template.EN_US,
			Components: components{
				Header: []string{"8967"},
				Body:   []string{"Ire", "8967", "15"},
			},
			Output: `{"name":"signup_otp_1","language":{"code":"en_US"},"components":[{"type":"header","parameters":[{"type":"text","text":"8967"}]},{"type":"body","parameters":[{"type":"text","text":"Ire"},{"type":"text","text":"8967"},{"type":"text","text":"15"}]}]}`,
		},
		{
			TemplateName: "signup_otp_1",
			Language:     template.EnglishUS,
			Components: components{
				Body: []string{"Ife", "8967", "15"},
			},
			Output: `{"name":"signup_otp_1","language":{"code":"en_US"},"components":[{"type":"body","parameters":[{"type":"text","text":"Ife"},{"type":"text","text":"8967"},{"type":"text","text":"15"}]}]}`,
		},
	}

	for _, tc := range tt {
		obj := template.New(tc.TemplateName, tc.Language)

		for _, h := range tc.Components.Header {
			obj.AddHeader(h)
		}
		for _, b := range tc.Components.Body {
			obj.AddBody(b)
		}
		for _, f := range tc.Components.Button {
			obj.AddButton(f)
		}

		got, er := obj.String()
		if er != nil {
			t.Errorf("error: %v", er)
		}

		if got != strings.TrimSpace(tc.Output) {
			t.Errorf("mismatched")
		}
	}

}

func TestCleanText(t *testing.T) {
	tests := []struct {
		Args     string
		Expected string
	}{
		{
			Args:     "Hi    	You doing?  ",
			Expected: "Hi You doing?",
		},
		{
			Args:     "Hi    	You doing?  ",
			Expected: "Hi You doing?",
		},
		{
			Args:     "Hi    	You doing?  ",
			Expected: "Hi You doing?",
		},
		{
			Args:     "Hi    	You doing?  ",
			Expected: "Hi You doing?",
		},
	}

	for i, tt := range tests {
		actual := template.CleanText(tt.Args)
		if actual != tt.Expected {
			t.Errorf(`Test (%d): expected "%s", actual "%s"`, i, tt.Expected, actual)
		}
	}
}
