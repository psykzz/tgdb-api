package main

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"os"
)

// Config - system config
type Config struct {
	Database struct {
		Host         string `yaml:"host" envconfig:"DB_HOST"`
		Port         string `yaml:"port" envconfig:"DB_PORT"`
		DatabaseName string `yaml:"name" envconfig:"DB_DBNAME"`
		Username     string `yaml:"user" envconfig:"DB_USERNAME"`
		Password     string `yaml:"pass" envconfig:"DB_PASSWORD"`
	} `yaml:"database"`
}

func processConfig(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		panic(err)
	}

	err = envconfig.Process("", cfg)
	if err != nil {
		panic(err)
	}
}
