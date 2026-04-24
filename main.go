package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nikit0xic/secretctl/auth"
	"gopkg.in/yaml.v3"
)

func main() {
	// if err := cmd.RootCmd.Execute(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	b := auth.Backend{}
	var configFilePath string = "/Users/nikitaterehov/PyCharmProjects/opensource/secretctl/auth/config.yaml"

	data, err := os.ReadFile(configFilePath)

	err = yaml.Unmarshal(data, &b)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", b)

	fmt.Println("Error: ", err)
	fmt.Println("Data: ", string(data))

	// yaml.Unmarshal([]byte(data), auth)

	a := auth.Backend{
		Name: "Vault",
		// Type: "vault",
		Address: "http://localhost:8200",
		Auth: auth.Auth{
			EnvVar: "VAULT_DEV_ROOT_TOKEN_ID",
			Exec: &auth.ExecAuth{
				Command: "cat",
				Args:    []string{"| grep", " -C", " 3"},
			},
		},
	}

	a.PrintBackend()
}
