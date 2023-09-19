package constants

var RegularExpression = struct {
	MOBILE    string
	NAME      string
	CardType  string
	Pincode   string
	ProductID string
	Action    string 
}{
	MOBILE:    "^\\d{10}$",
	NAME:      "^[a-zA-Z0-9./\\-,_() ']{1,70}$",
	CardType:  "^(credit|debit)$",
	Pincode:   "^\\d{6}$",
	ProductID: "^[a-zA-Z0-9]{1,70}$",
	Action:    "^(add|delete)$",
}
