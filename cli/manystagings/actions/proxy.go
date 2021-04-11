package actions

import "github.com/urfave/cli"

func (a *Actions) ProxyApp(c *cli.Context) error {
	return a.orchestratorCLI.ProxyApp("carlos", "nginx")
}
