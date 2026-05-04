package vault

import (
	"fmt"
	"os/exec"

	"github.com/nikit0xic/secretctl/auth"
)

func ConnectVaultBackend(b auth.Backend) {
	cmd := exec.Command("vault", "kv", "list", "-address="+b.Address, "secret/")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(out))
}

func GetCmd(b auth.Backend, args []string) {

	for i := 0; i < len(args); i++ {
		vault_exec := exec.Command("vault", "kv", "get", "-address="+b.Address, args[i])
		out, err := vault_exec.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
		}
		fmt.Println(string(out))
	}
}
