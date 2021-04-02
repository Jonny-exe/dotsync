package conf

import (
	// . "github.com/WAY29/icecream-go/icecream"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type conf struct {
	AccessToken string `yaml:"gh-access-token"`
	GhUsername  string `yaml:"gh-username"`
}

func getConf() conf {
	var c conf
	home := os.Getenv("HOME")
	yamlFile, err := ioutil.ReadFile(home + "/.dotsync.yaml")
	if err != nil {
		log.Error("Error reading config")
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Error("Config file syntax is incorrect, or has something missing")
	}
	return c
}

// Conf is the user configuration
var Conf conf = getConf()
