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
		Long:  `A Fast and Flexible staging manager built in Go.`,
	}

	// Configure
	rootCmd.AddCommand(&cobra.Command{
		Use:   "configure",
		Short: "configure the manystagins CLI",
		Run: func(cmd *cobra.Command, args []string) {
			a.Configure()
		},
	})

	// Up
	var upRecreate bool
	upCmd := &cobra.Command{
		Use:   "up",
		Short: "up all or an application for staging",
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.Up(args, upRecreate)
		},
	}
	rootCmd.AddCommand(upCmd)
	upCmd.Flags().BoolVarP(&upRecreate, "recreate", "r", false, "recreate app if already exists")

	// Kill
	rootCmd.AddCommand(&cobra.Command{
		Use:   "kill",
		Short: "kill all or an application for staging",
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.Kill(args)
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
		Use:     "exec",
		Short:   "execute a command into an application container",
		Example: "manystagings exec [APP_NAME] -- [COMMAND]",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("application name required")
			}
			if cmd.ArgsLenAtDash() <= 0 {
				return errors.New("missing exec command. Please run:\nmanystagings exec [APP_NAME] -- [COMMAND]")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			argsLenAtDash := cmd.ArgsLenAtDash()
			execArgs := args[argsLenAtDash:]
			return a.ExecDeployment(args[0], execArgs)
		},
	})

	// Status
	rootCmd.AddCommand(&cobra.Command{
		Use:     "status",
		Aliases: []string{"ps"},
		Short:   "get application statuses inside your staging",
		Example: "manystagings status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.Status()
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
