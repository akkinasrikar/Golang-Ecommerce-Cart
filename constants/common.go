package constants

var ProductConstants = struct {
	ADDITION      string
	DELETION      string
	SUCCESS       string
	FAILED        string
	CARD          string
	WALLET        string
	CARDANDWALLET string
	ONTIME        string
	DELIVERED     string
}{
	ADDITION:      "add",
	DELETION:      "delete",
	SUCCESS:       "success",
	FAILED:        "failed",
	CARD:          "card",
	WALLET:        "wallet",
	CARDANDWALLET: "card & wallet",
	ONTIME:        "on time",
	DELIVERED:     "delivered",
}

var ProcessTasks = struct {
	IMAGERESIZE string
	SENDEMAIL   string
}{
	IMAGERESIZE: "image-resize",
	SENDEMAIL:   "send-email",
}
