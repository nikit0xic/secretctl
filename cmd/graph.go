package cmd

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"

	"github.com/nikit0xic/secretctl/auth"
	// "github.com/nikit0xic/secretctl/gitlab"
	"github.com/nikit0xic/secretctl/vault"
)

var graphCmd = &cobra.Command{

	Use:   "graph",
	Short: "Get graph of secrets from path",
	Args:  cobra.ExactArgs(1),
	Run:   runGraph,
}

func runGraph(cmd *cobra.Command, args []string) {
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

	if len(args) < 1 {
		fmt.Println("Error: path argument is required")
		return
	}

	depth, _ := cmd.Flags().GetInt("level")

	for i := range backends {

		switch backends[i].Type {
		case "vault":
			sv := vault.NewSafeVisited(backends[i])
			var wg sync.WaitGroup
			wg.Add(1)
			go vault.GetGraph(backends[i], args[0], depth, 0, sv, &wg)
			wg.Wait()

			if err != nil {
				fmt.Print("Error: ", err)
				continue
			}
		case "gitlab":
			fmt.Println("You've triggeed gitlab list command. Yet this command is not complete for use.")
		default:
			fmt.Println("Your contexts backends not supported for get command yet!")
		}
	}
}
