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
	Kind     string `json:"kind"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	DBName   string `json:"dbName"`
	SslMode  string `json:"sslMode"`
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
