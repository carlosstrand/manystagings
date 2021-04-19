package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/carlosstrand/manystagings/cli/manystagings/actions"
	"github.com/carlosstrand/manystagings/cli/manystagings/client"
	"github.com/carlosstrand/manystagings/cli/manystagings/orchestratorcli"
	"github.com/carlosstrand/manystagings/cli/manystagings/orchestratorcli/providerscli/kubernetescli"
	"github.com/carlosstrand/manystagings/cli/manystagings/utils/msconfig"
	"github.com/spf13/cobra"
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
	client := client.NewClient(config.HostURL)
	client.SetAuthToken(config.Token)
	a := actions.NewActions(actions.Options{
		OrchestratorCLI: orchestratorCLI,
		Client:          client,
		Config:          config,
	})

	var rootCmd = &cobra.Command{
		Use:   "manystagings",
		Short: "Setup your staging environment easily with manystagings",
		Long:  `A Fast and Flexible staging manager Site Generator built in Go.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}

	// Configure
	rootCmd.AddCommand(&cobra.Command{
		Use:   "configure",
		Short: "configure the manystagins CLI",
		Run: func(cmd *cobra.Command, args []string) {
			a.Configure()
		},
	})

	// Proxy
	rootCmd.AddCommand(&cobra.Command{
		Use:   "proxy",
		Short: "proxy to an application inside your staging",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("application name required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.ProxyDeployment(args[0])
		},
	})

	// Exec
	rootCmd.AddCommand(&cobra.Command{
		Use:   "exec",
		Short: "execute a command into an application container",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("application name required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.ExecDeployment(args[0])
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
