package cmd

type PluginHandler interface {
	Lookup(filename string) (string, bool)

	Execute(executePath string, cmdArgs, environment []string) error
}
