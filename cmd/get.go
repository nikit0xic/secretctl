package cmd

import (
	"fmt"

	"github.com/nikit0xic/secretctl/auth"
	"github.com/nikit0xic/secretctl/gitlab"
	"github.com/nikit0xic/secretctl/vault"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{

	Use:   "get",
	Short: "Get secrets from configured backends",
	Run:   runGetCmd,
}

func runGetCmd(cmd *cobra.Command, args []string) {
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
			vault.GetCmd(b[i], args)
		case "gitlab":
			gitlab.GetCmd()
		default:
			fmt.Println("Your contexts backends not supported for get command yet!")
		}
	}
}
