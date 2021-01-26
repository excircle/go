package main

import (
	"fmt"
	"log"
	"os"

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
	var yamlConfig Config

	//Populate yamlConfigData - Feel free to change the data
	yamlConfig.Domain.Hostname = "example.com"
	yamlConfig.Domain.SOA = []string{"ns11.example.com", "ns12.example.com"}
	yamlConfig.Domain.IP = "10.0.0.80"
	yamlConfig.Domain.Nameservers = []string{"ns11.example.com", "ns12.example.com"}

	//Create []bytes from yaml struct 'yamlConfig'
	sliceOfBytes, err := yaml.Marshal(&yamlConfig)
	if err != nil {
		log.Fatalf("We could not Marshal new struct: %s\n", err)
	}

	////create new YAML config 'b.yaml'
	f, err := os.Create("./b.yaml")
	if err != nil {
		log.Fatalf("We could not Create new file 'b.yaml': %s\n", err)
	}

	//If anything fails, execute a close command no matter what
	defer f.Close()

	////write to b.yaml
	f.Write(sliceOfBytes)

	fmt.Println("Please examine the contents of 'b.yaml' file for correctness")
}
