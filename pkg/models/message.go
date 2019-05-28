package models

//easyjson:json
type Message struct {
	Body      string `json:"body"`
	Nickname  string `json:"nickname"`
	Timestamp string `json:"timestamp"`
	Edited    bool   `json:"edited"`
	Imgurl    string `json:"imgurl"`
}
