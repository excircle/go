package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
)

func main() {
	// Example values for your Vault env
	vault_addr := "http://127.0.0.1:8200"
	token := "ROOT_TOKEN_VALUE"

	// Create a api.Config struct with specific vaules
	config := &api.Config{
		Address: vault_addr,
	}

	// Create new client struct
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add Vault's root token to client
	client.SetToken(token)

	// Obtain Vault init status
	initStatus, err := client.Sys().InitStatus()
	if err != nil {
		fmt.Println(err)
		return
	}

	//Print returns "True" if Vault is successfully initialized
	fmt.Println(initStatus)
}
