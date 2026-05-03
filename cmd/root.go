package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nikit0xic/secretctl/auth"
	"github.com/nikit0xic/secretctl/gitlab"
	"github.com/nikit0xic/secretctl/vault"
)

var (
	uppercase bool
)

var RootCmd = &cobra.Command{
	Use:     "secretctl",
	Aliases: []string{"sectl"},
	Short:   "A tool for secret management for platforms such as Vault, GitLab etc.",

	Args: cobra.ArbitraryArgs,
	Run:  runSecretctlCmd,
}

func init() {
	RootCmd.Flags().BoolVarP(&uppercase, "env", "e", false, "Env from flag")
}

func runSecretctlCmd(cmd *cobra.Command, args []string) {
	cfg, _ := auth.LoadConfig("")

	Backends, error := cfg.GetBackendsForContext(cfg.CurrentContext)

	if error != nil {
		fmt.Print("Error: ", error)
	}

	fmt.Println("Current context:", cfg.CurrentContext)

	for i := range Backends {
		switch Backends[i].Type {
		case "vault":
			vault.ConnectVaultBackend()
		case "gitlab":
			gitlab.ConnectGitlabBackend()
		default:
			fmt.Println("Your contexts backends not supported yet!")
		}
	}
	for i, _ := range Backends {
		fmt.Println("Backend #", i, "is: ", Backends[i])
	}
}
