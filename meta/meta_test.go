package meta_test

import (
	"github.com/iamNator/go-whatsapp/meta"
	"strings"
	"testing"
)

func TestBytes(t *testing.T) {
	obj := meta.New("signup_otp_1", "+2349045057268", "en_US")

	obj1 := obj.
		AddHeader("8967").
		AddBody("Ire").
		AddBody("8967").
		AddBody("15")

	b, er := obj1.Byte()
	if er != nil {
		t.Errorf("error: %v", er)
	}

	if string(b) != strings.TrimSpace(`{"messaging_product":"whatsapp","to":"2349045057268","type":"template","template":{"name":"signup_otp_1","language":{"code":"en_US"},"components":[{"type":"header","parameters":[{"type":"text","text":"8967"}]},{"type":"body","parameters":[{"type":"text","text":"Ire"},{"type":"text","text":"8967"},{"type":"text","text":"15"}]}]}}`) {
		t.Errorf("mismatched")
	}

}
