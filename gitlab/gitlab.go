package gitlab

import (
	"fmt"
)

func ConnectGitlabBackend() {
	fmt.Println("You triggered gitlab backend")
}

func GetCmd() {
	fmt.Println("You triggered gitlab backend's get command")
}
