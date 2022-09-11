package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

func Load() *Config {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	c := new(Config)
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
