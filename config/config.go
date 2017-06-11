package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	DbUrl      string
	Darkskyapi string
	MyUrl      string
}

func GetConfig() (conf Configuration, err error) {
	file, err := os.Open("conf.json")
	if err != nil {
		return
	}
	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		return
	}
	return
}
