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
	CurCtx := cfg.CurrentContext
	var Ctx auth.Context
	for _, v := range cfg.Contexts {
		if v.Name == CurCtx {
			Ctx = v
		}
	}

	backs := Ctx.Backends

	fmt.Println("Current context:", cfg.CurrentContext)
	for i, v := range backs {
		if v == "vault" {
			vault_exec := exec.Command("vault", "kv", "list", "/secret")
			out, err := vault_exec.CombinedOutput()
			if err != nil {
				fmt.Errorf("Error: ", err)
			}
			defer fmt.Println(string(out))
		}
		fmt.Println("Backend #", i, "is: ", backs[i])
	}

}
