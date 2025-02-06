package config

import (
	"encoding/json"
	"log"
	"os"
)

type PgConfig struct {
	Host string 		`json:"host"`
	Port int    		`json:"port"`
	User string 		`json:"user"`
	Password string `json:"password"`
	DBname string 	`json:"dbname"`
}

type Configuration struct {
	NinjaApiKey  string `json:"ninja-api-key"`
	PG PgConfig `json:"pg"`
}

var Secrets *Configuration = &Configuration{}

// todo make this generic config loader when needed
func LoadConfig (filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal("unable to open file configuration file: ", filepath)
		os.Exit(1)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(Secrets); err != nil {
		log.Fatal("error parsing configuration file: ", err)
		os.Exit(1)
	}
}