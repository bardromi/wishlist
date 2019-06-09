package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	WebConfiguration      `json:"webConfiguration"`
	DatabaseConfiguration `json:"databaseConfiguration"`
}

type WebConfiguration struct {
	Address string `json:"address"`
}

type DatabaseConfiguration struct {
	User       string `json:"username"`
	Password   string `json:"password"`
	Host       string `json:"address"`
	Name       string `json:"dbName"`
	DisableTLS bool   `json:"sslMode"`
}

func LoadConfig() *Configuration {
	jsonFile, err := os.Open("C:/Users/Bar/go/src/github.com/bardromi/wishlist/cmd/server/config.json")
	if err != nil {
		log.Fatalln("Cannot open config file", err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	conf := Configuration{}

	err = json.Unmarshal(byteValue, &conf)

	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}

	return &conf
}
