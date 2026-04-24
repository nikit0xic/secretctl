package auth

import (
	"fmt"
)

type Config struct {
	CurrentContext string `yaml:"current-context"`

	Contexts []Context `yaml:"contexts"`
	Backends []Backend `yaml:"backends"`
}

type Context struct {
	Name    string  `yaml:"name"`
	Backend Backend `yaml:"backend"`
}

// TODO: Type typeOfBackend: aws, vault, gitlab ...
type Backend struct {
	Name string `yaml:"name"`
	// Type    string
	Address string `yaml:"address"`
	Auth    Auth   `yaml:"auth"`
}

type Auth struct {
	EnvVar string    `yaml:"env,omitempty"`
	Exec   *ExecAuth `yaml:"exec,omitempty"`
}

type ExecAuth struct {
	Command string   `yaml:"command"`
	Args    []string `yaml:"args"`
}

func (B Backend) PrintBackend() {
	fmt.Println(B.Name)
	fmt.Println(B.Address)
	fmt.Println(B.Auth)
}
