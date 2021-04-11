package main

import (
	"fmt"
	"log"
	"os"

	"github.com/carlosstrand/manystagings/cli/manystagings/actions"
	"github.com/carlosstrand/manystagings/cli/manystagings/orchestratorcli"
	"github.com/carlosstrand/manystagings/cli/manystagings/orchestratorcli/providerscli/kubernetescli"
	"github.com/carlosstrand/manystagings/cli/manystagings/utils/msconfig"
	"github.com/urfave/cli"
)

func main() {
	config, err := msconfig.LoadConfig()
	if err != nil {
		fmt.Println("Could not load config. Please run:\n\n\t manystagings configure")
		os.Exit(1)
	}
	var orchestratorCLI orchestratorcli.OrchestratorCLI
	switch config.OrchestratorProvider {
	case "kubernetes":
		orchestratorCLI = kubernetescli.NewKubernetesCLIProvider(kubernetescli.Options{
			Config: config,
		})
	}
	a := actions.NewActions(actions.Options{
		OrchestratorCLI: orchestratorCLI,
	})
	app := &cli.App{
		Name:  "configure",
		Usage: "configure the manystagins CLI",
		Commands: []cli.Command{
			{
				Name:    "configure",
				Aliases: []string{"c"},
				Usage:   "configure the manystagins CLI",
				Action: func(c *cli.Context) error {
					return a.Configure()
				},
			},
			{
				Name:    "proxy",
				Aliases: []string{"p"},
				Usage:   "configure the manystagins CLI",
				Action: func(c *cli.Context) error {
					return a.ProxyApp(c)
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
