package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/nikit0xic/secretctl/auth"
)

var (
	uppercase bool
)

var RootCmd = &cobra.Command{
	Use:     "secretctl",
	Aliases: []string{"sectl"},
	Short:   "A tool for secret management for platforms such as Vault, GitLab etc.",

	Args: cobra.ArbitraryArgs,
	Run:  runTextFormatter,
}

func init() {
	RootCmd.Flags().BoolVarP(&uppercase, "upper", "u", false, "Convert text to uppercase")
}

func runTextFormatter(cmd *cobra.Command, args []string) {
	var configFilePath string = "/Users/nikitaterehov/PyCharmProjects/opensource/secretctl/auth/config.yaml"

	data, err := os.ReadFile(configFilePath)

	if err != nil {
		fmt.Println("Read error: %v", err)
	}

	cfg := auth.Config{}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("Config struct:\n%v\n\n", cfg)
}
