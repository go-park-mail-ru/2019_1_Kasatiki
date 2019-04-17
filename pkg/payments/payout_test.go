package payments

import (
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/models"
	"testing"
)

func TestPhonePayBody(t *testing.T) {
	tested := `{
        "id":"1",
        "sum": {
          "amount":1,
          "currency":"643"
        },
        "paymentMethod": {
          "type":"Account",
          "accountId":"643"
        },
        "fields": {
          "account":"1"
        }
      }`
	body := PhonePayBody("1", "1", "1")
	if tested != body {
		t.Errorf("Not valid phone payments body generated")
	}

}

func TestReadConfig(t *testing.T) {
	var testedBill models.Credentials
	err := ReadConfig("sdfweq", &testedBill)
	if err == nil {
		t.Errorf("Expected nil")
	}
	err = ReadConfig("./resources/credentials.json", &testedBill)
	if testedBill.Wallet != "79851301115" {
		t.Errorf("Not valid decode")
	}
}
