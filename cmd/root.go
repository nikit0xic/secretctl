package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

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
	Run:  runSecretctlCmd,
}

func init() {
	RootCmd.Flags().BoolVarP(&uppercase, "env", "e", false, "Env from flag")
}

func runSecretctlCmd(cmd *cobra.Command, args []string) {
	fmt.Println(auth.LoadConfig(""))
}
