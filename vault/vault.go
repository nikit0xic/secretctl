package vault

import (
	"fmt"
	"os/exec"
)

func ConnectVaultBackend() {
	vault_exec := exec.Command("vault", "kv", "list", "/secret")
	out, err := vault_exec.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(out))
}
