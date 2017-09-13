package config

import (
	"encoding/json"
	"errors"
	"os"
)

type Configuration struct {
	ListenAddress  int
	FirebaseConfig FirebaseConfig
	SigningKey     string
}

type FirebaseConfig struct {
	URl string
}

var (
	ErrCantFindConfig = errors.New("Config file is missing")
)

func NewConfig(configfile string) (*Configuration, error) {

	if _, err := os.Stat(configfile); os.IsNotExist(err) {
		return nil, ErrCantFindConfig
	}

	configuration := Configuration{}

	file, _ := os.Open(configfile)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}

	return &configuration, nil

}

func MustNewConfig(configfile string) *Configuration {
	config, err := NewConfig(configfile)

	if err != nil {
		panic(err)
	}

	return config

}
