package parsers

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Timeout int
	Servers []struct {
		Name string
		Url string
	}
}

func ParseConfig() Config {
	var conf Config
	configFileName := "config.yaml"

	log.Println("Parsing configuration file")

	configFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Println("Config file not found, error:", err)
	}
	err = yaml.Unmarshal(configFile, &conf)
	if err != nil {
		log.Println("Error unmarshalling configuration file:", err)
	}

	return conf
}
