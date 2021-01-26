package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Config
type Config struct {
	Domain struct {
		Hostname    string   `yaml:"hostname"`
		SOA         []string `yaml:"soa,flow"`
		IP          string   `yaml:"ip"`
		Nameservers []string `yaml:"nameservers,flow"`
	}
}

func main() {
	//Parse options
	configFile := "./config.yaml"
	var yamlConfig Config

	//Read configFile into a slice of bytes ([]byte)
	sliceOfBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
	}

	//Unmarshal bytes into 'yamlConfig' (Type == Config)
	err = yaml.Unmarshal(sliceOfBytes, &yamlConfig)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	}

	fmt.Println(yamlConfig)
}
