package payments

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	//"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/models"
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

type OperatorsInfo struct {
	Message             string `json:"message"`
	Code 				*OperatorStatus  `json:"code"`
}

type OperatorStatus struct {
	Value string `json:"value"`
	Name string `json:"_name"`
}



func CheckOperator(phoneNumber string) (string, error){
	var jsonStr = []byte(`phone=7` + phoneNumber)
	fmt.Println(bytes.NewBuffer(jsonStr))
	requestToDetect, err :=  http.NewRequest("POST", "https://qiwi.com/mobile/detect.action", bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	requestToDetect.Header.Add("Accept", "application/json")
	requestToDetect.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(requestToDetect)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)

	var info OperatorsInfo

	_ = json.Unmarshal([]byte(body), &info)
	fmt.Println(info.Message)
	if info.Code.Name == "NORMAL" {
		return info.Message, nil
	} else {
		return "", errors.New("Failed recognition")
	}
}




func PhonePayout(paymentsInfo models.Credentials, inputPhone string, amount string) error {
	err := ReadConfig("credentials.json",
		&paymentsInfo)
	if err != nil {
		fmt.Println(err)
		return errors.New("Unable to get last transaction id")
	}

	timestamp := time.Now().Unix()
	trnid := strconv.Itoa(int(1000 * timestamp))
	fmt.Print(trnid)

	// Check class
	oper, err := CheckOperator(inputPhone)
	if err != nil {
		return err
	}
	ppRequest, err := http.NewRequest("POST", "https://edge.qiwi.com/sinap/api/v2/terms/" + oper + "/payments",
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
