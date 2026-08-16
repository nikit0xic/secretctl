package vault

import (
	"fmt"
	// "go/printer"
	"os/exec"
	"strings"
	"sync"

	"github.com/nikit0xic/secretctl/auth"
)

// https://gist.github.com/GeorgeHernandez/10dcbb5fd6ca8b087d169d5a44d72cd2
const (
	horizontal_connector = '─'
	vertical_connector   = '|'
	branch_connector     = '├'
	leaf_connector       = '└'
)

type SafeVisited struct {
	mu      sync.Mutex
	backend auth.Backend
	visited map[string]bool
}

type Fetcher interface {
	List(b auth.Backend, path string) (s []string, err error)
}

func NewSafeVisited(b auth.Backend) *SafeVisited {
	return &SafeVisited{
		backend: b,
		visited: map[string]bool{},
	}
}

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

func GetGraph(backend auth.Backend, path string, depth int, currentDepth int, sv *SafeVisited, wg *sync.WaitGroup) {
	defer wg.Done()

	sv.mu.Lock()
	if sv.visited[path] {
		sv.mu.Unlock()
		return
	}

	if depth != -1 && currentDepth >= depth {
		return
	}

	sv.visited[path] = true
	sv.mu.Unlock()

	entries, err := GetList(backend, []string{path})
	if err != nil {
		return
	}

	for i, entry := range entries {
		indent := strings.Repeat("  ", currentDepth)
		isLast := i == len(entries)-1
		connector := string(branch_connector)

		if isLast {
			connector = string(leaf_connector)
		}

		fmt.Println(indent + connector + "── " + entry)

		if strings.HasSuffix(entry, "/") {
			wg.Add(1)
			go GetGraph(backend, path+entry, depth, currentDepth+1, sv, wg)
		}
	}

}

func GetList(b auth.Backend, args []string) ([]string, error) {
	var path string
	if len(args) == 0 {
		path = "secret/"
	} else {
		path = args[0]
	}
	cmd := exec.Command("vault", "kv", "list", "--address="+b.Address, path)
	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	lines := strings.Split(string(out), "\n")
	var payload []string

	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "Keys") || strings.HasPrefix(line, "---") {
			continue
		}
		payload = append(payload, line)
	}

	return payload, nil
}
