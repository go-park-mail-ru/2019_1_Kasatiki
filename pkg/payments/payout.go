package payments

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func PhonePayBody(trnid string, phone string, amount string) string {

	body := `{
        "id":"` + trnid + `",
        "sum": {
          "amount":` + amount + `,
          "currency":"643"
        },
        "paymentMethod": {
          "type":"Account",
          "accountId":"643"
        },
        "fields": {
          "account":"` + phone + `"
        }
      }`
	fmt.Println(body)
	return body
}

func ReadConfig(path string, paymentsInfo *models.Credentials) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New("Unable to read file from path:" + path)
	}
	err = json.Unmarshal([]byte(file), paymentsInfo)
	if err != nil {
		return errors.New("Wrong json format of credentials file")
	}
	return nil
}

func PhonePayout(paymentsInfo models.Credentials, inputPhone string, amount string) error {
	err := ReadConfig("/home/evv/GolandProjects/2019_1_Kasatiki/pkg/payments/resources/credentials.json",
		&paymentsInfo)
	if err != nil {
		return errors.New("Unable to get last transaction id")
	}

	timestamp := time.Now().Unix()
	trnid := strconv.Itoa(int(1000 * timestamp))
	fmt.Print(trnid)

	ppRequest, err := http.NewRequest("POST", "https://edge.qiwi.com/sinap/api/v2/terms/1/payments",
		strings.NewReader(PhonePayBody(trnid, inputPhone, amount)))
	ppRequest.Header.Add("Accept", "application/json")
	ppRequest.Header.Add("Content-Type", " application/json")
	ppRequest.Header.Add("Authorization", "Bearer "+paymentsInfo.Token)

	client := http.Client{}
	resp, err := client.Do(ppRequest)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	readableBody := string(bytes)
	defer resp.Body.Close()
	fmt.Println(readableBody)

	return nil
}
