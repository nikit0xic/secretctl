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

	b, err := cfg.GetBackendsForContext(cfg.CurrentContext)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i := range b {
		switch b[i].Type {
		case "vault":
			vault.GetList(b[i], args)
		case "gitlab":
			fmt.Println("You've triggeed gitlab list command. Yet this command is not complete for use.")
		default:
			fmt.Println("Your contexts backends not supported for get command yet!")
		}
	}
}
