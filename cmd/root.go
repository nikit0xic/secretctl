package cmd

import (
	"fmt"
	"os/exec"

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
	cfg, _ := auth.LoadConfig("")
	backs := cfg.Backends

	for i, _ := range backs {
		if backs[i].Name == "vault" {
			vault_exec := exec.Command("vault", "kv", "list", "/secret")
			out, err := vault_exec.CombinedOutput()
			if err != nil {
				fmt.Errorf("Error: ", err)
			}
			fmt.Println(string(out))
		}
		fmt.Println("the 'back of' i:", i, "is: ", backs[i])
	}

}
