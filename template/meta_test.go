package template_test

import (
	"strings"
	"testing"

	"github.com/iamNator/go-whatsapp/template"
)

func TestBytes(t *testing.T) {
	obj := template.New("signup_otp_1", "+2349045057268", "en_US")

	b, er := obj.
		AddHeader("8967").
		AddBody("Ire").
		AddBody("8967").
		AddBody("15").
		Byte()
	if er != nil {
		t.Errorf("error: %v", er)
	}

	if string(b) != strings.TrimSpace(`{"messaging_product":"whatsapp","to":"2349045057268","type":"template","template":{"name":"signup_otp_1","language":{"code":"en_US"},"components":[{"type":"header","parameters":[{"type":"text","text":"8967"}]},{"type":"body","parameters":[{"type":"text","text":"Ire"},{"type":"text","text":"8967"},{"type":"text","text":"15"}]}]}}`) {
		t.Errorf("mismatched")
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