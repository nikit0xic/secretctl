package auth

import (
	"fmt"
)

type Config struct {
	CurrentContext string    `yaml:"current-context"`
	Contexts       []Context `yaml:"contexts"`
	Backends       []Backend `yaml:"backends"`
}

type Context struct {
	Name     string   `yaml:"name"`
	Backends []string `yaml:"backend"`
}

type Backend struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
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

func (cfg *Config) GetBackendsForContext(contextName string) ([]Backend, error) {
	var ctx *Context
	for _, v := range cfg.Contexts {
		if v.Name == contextName {
			ctx = &v
		}
	}

	if ctx == nil {
		return nil, fmt.Errorf("context '%s' not found", contextName)
	}

	var backs []Backend

	for _, backendName := range ctx.Backends {
		for _, b := range cfg.Backends {
			if backendName == b.Name {
				backs = append(backs, b)
			}
		}
	}

	return backs, nil
}
