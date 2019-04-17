package payments

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2019_1_Kasatiki/pkg/models"
	"io/ioutil"
)

func ReadConfig(path string, paymentsInfo models.Credentials) error {
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

func GetLastTransactionId(paymentsInfo models.Credentials) error {
	err := ReadConfig("/home/evv/GolandProjects/2019_1_Kasatiki/pkg/payments/resources/credentials.json",
		paymentsInfo)
	if err != nil {
		return errors.New("Unable to get last transaction id")
	}

	//url := fmt.Sprintf(paymentsInfo.TransactionInfo, paymentsInfo.Wallet)
	fmt.Println(paymentsInfo.TransactionInfo, paymentsInfo.Wallet)
	//req, err := http.NewRequest("GET", url, nil)
	//req.Header.Add("Accept", "application/json")
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Authorization", "Bearer"+paymentsInfo.Token)
	//
	//if err != nil {
	//	return errors.New("Unable to prepare post request")
	//}
	//client := http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	return err
	//}
	//defer resp.Body.Close()
	//bytes, err := ioutil.ReadAll(resp.Body)
	//readableBody := string(bytes)
	//fmt.Println(readableBody)
	return nil
}
