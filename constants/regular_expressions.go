package constants

var RegularExpression = struct {
	MOBILE   string
	NAME     string
	CardType string
	Pincode  string
}{
	MOBILE:   "^\\d{10}$",
	NAME:     "^[a-zA-Z0-9./\\-,_() ']{1,70}$",
	CardType: "^(credit|debit)$",
	Pincode:  "^\\d{6}$",
}
