package models

//easyjson:json
type Config struct {
	DBHost     string `json:"DatabaseHost"`
	DBPort     uint16 `json:"DatabasePort"`
	DBUser     string `json:"DatabaseUser"`
	DBPassword string `json:"DatabasePassword"`
	DBSpace    string `json:"DatabaseSpace"`
	ServerPort int    `json:"ServerPort"`
}
