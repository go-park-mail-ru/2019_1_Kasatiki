package models

//easyjson:json
type Payout struct {
	Amount string `json:"amount"`
	Phone  string `json:"phone"`
}

//easyjson:json
type Credentials struct {
	Wallet             string `json:"wallet"`
	Token              string `json:"token"`
	TransactionInfo    string `json:"lasttn"`
	PaymentVisa        string `json:"visa"`
	PaymentsMasterCard string `json:"mastercard"`
	PaymentPhone       string `json:"phone"`
}
