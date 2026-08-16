package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nikit0xic/secretctl/auth"
	"github.com/nikit0xic/secretctl/vault"
)

var listCmd = &cobra.Command{

	Use:   "list",
	Short: "Get list of secrets in specific path",
	Args:  cobra.MaximumNArgs(1),
	Run:   runListCmd,
}

func runListCmd(cmd *cobra.Command, args []string) {
	cfg, err := auth.LoadConfig("")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	backends, err := cfg.GetBackendsForContext(cfg.CurrentContext)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i := range backends {
		switch backends[i].Type {
		case "vault":
			payload, err := vault.GetList(backends[i], args)
			if err != nil {
				fmt.Print("Error: ", err)
				continue
			}
			fmt.Printf("=== Backend: %s ===\n", backends[i].Address)
			for _, p := range payload {
				fmt.Println(p)
			}
		case "gitlab":
			fmt.Println("You've triggeed gitlab list command. Yet this command is not complete for use.")
		default:
			fmt.Println("Your contexts backends not supported for get command yet!")
		}
	}
}
