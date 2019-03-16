package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type credentials struct {
	Wallet             string `json:"wallet"`
	Token              string `json:"token"`
	TransactionInfo    string `json:"lasttn"`
	PaymentVisa        string `json:"visa"`
	PaymentsMasterCard string `json:"mastercard"`
}

type payout struct {
	Phone  string `json:"phone"`
	Amount string `json:"amount"`
}

var paymentsInfo credentials

func init() {
	err := paymentsInfo.intitialize("/home/evv/GolandProjects/2019_1_Kasatiki/payments/resources/credentials.json")
	if err != nil {
		log.Fatal("Cant read credentials for payout systems: ", err)
	}
}
func getLastTransactionId() error {
	err := paymentsInfo.intitialize(paymentsInfo.TransactionInfo)
	if err != nil {
		return errors.New("Unable to get last transaction id")
	}

	url := fmt.Sprintf(paymentsInfo.TransactionInfo, paymentsInfo.Wallet)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer"+paymentsInfo.Token)

	if err != nil {
		return errors.New("Unable to prepare post request")
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	readableBody := string(bytes)
	fmt.Println(readableBody)
	return nil
}

func readConfig(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.New("Unable to read file from path:" + path)
	}
	err = json.Unmarshal([]byte(file), &paymentsInfo)
	if err != nil {
		return errors.New("Wrong json format of credentials file")
	}
	return nil
}

func (info *credentials) intitialize(path string) error {
	_ = readConfig("/home/evv/GolandProjects/2019_1_Kasatiki/payments/resources/credentials.json")
	return nil
}

//func (urls *pUrls) initialize (path string) error {
//
//}
func main() {
	//err := paymentsInfo.intitialize("/home/evv/GolandProjects/2019_1_Kasatiki/payments/resources/credentials.json")
	fmt.Println(paymentsInfo)
	err := getLastTransactionId()
	fmt.Println(err)
	//fmt.Println(err)
}
