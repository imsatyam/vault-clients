// simple client to interact with hashicorp vault
package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"
)

func main() {

	token := os.Getenv("VAULT_TOKEN")
	vaultUri := os.Getenv("VAULT_ADDR")

	client, err := api.NewClient(&api.Config{
		Address: vaultUri,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	client.SetToken(token)
	c := client.Logical()
	secretValues, err := c.Read("secret/vaultdemo")
	if err != nil {
		fmt.Println(err)
		return
	}
	for propName, propValue := range secretValues.Data {
		fmt.Printf(" - %s -> %v\n", propName, propValue)
	}
}
