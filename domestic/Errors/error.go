package Errors

// map of maps(JSONs, which i send on frontend)
var Error = map[string]map[string]string{
	"User does not exist":    {"Error": "User does not exist"},
	"Invalid password":       {"Error": "Invalid password"},
	"Nickname already exist": {"Error": "Nickname already exist"},
	"Email already exist":    {"Error": "Email already exist"},
	"Invalid JSON":           {"Error": "Invalid JSON"},
	"Not authorized":         {"Error": "Not authorized"},
	"Bad cookies":            {"Error": "Bad cookies"},
	"Bad File":               {"Error": "Bad File"},
}
