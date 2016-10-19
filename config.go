package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type instanceConfig struct {
    Hostname string
    Username string
    Password string
    Database string
    Driver   string
    Port     string
    ServerPort string
}

func (config *instanceConfig) readConfiguration() {
	configFile, _ := filepath.Abs("./config.yml")
	yamlFile, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value:  %#v\n", config)
}
