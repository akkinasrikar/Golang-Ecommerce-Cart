package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type FakeStoreConfig struct {
	BaseUrl string `json:"FAKESTORE_BASEURL"`
}

var FakeStore *FakeStoreConfig

func loadFakeStoreConfig() {
	FakeStore = &FakeStoreConfig{}
	err := envconfig.Process("fakeStore", FakeStore)
	if err != nil {
		log.Fatal(err.Error())
	}
}
