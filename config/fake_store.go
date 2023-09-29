package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type FakeStoreConfig struct {
	BaseUrl      string `json:"FAKESTORE_BASEURL"`
	PublicKey    string `json:"FAKESTORE_PUBLICKEY"`
	PrivateKey   string `json:"FAKESTORE_PRIVATEKEY"`
	Gmail        string `json:"FAKESTORE_GMAIL"`
	MailPassword string `json:"FAKESTORE_MAILPASSWORD"`
}

var FakeStore *FakeStoreConfig

func loadFakeStoreConfig() {
	FakeStore = &FakeStoreConfig{}
	err := envconfig.Process("fakeStore", FakeStore)
	if err != nil {
		log.Fatal(err.Error())
	}
}
