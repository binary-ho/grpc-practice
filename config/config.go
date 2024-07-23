package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Paseto struct {
		Key string
	}

	GRPC struct {
		URL string
	}
}

func NewConfig(path string) *Config {
	file := getFile(path)
	defer file.Close()

	config := new(Config)
	if err := toml.NewDecoder(file).Decode(config); err != nil {
		panic(err)
	}
	return config
}

func getFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}
