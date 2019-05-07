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

	config := &api.Config{
		Address: vaultUri,
	}
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	client.SetToken(token)
	c := client.Logical()
	secret, err := c.Read("secret/vaultdemo")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(secret)

}
