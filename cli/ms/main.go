package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/carlosstrand/manystagings/cli/ms/actions"
	"github.com/carlosstrand/manystagings/cli/ms/client"
	"github.com/carlosstrand/manystagings/cli/ms/orchestratorcli"
	"github.com/carlosstrand/manystagings/cli/ms/orchestratorcli/providerscli/kubernetescli"
	"github.com/carlosstrand/manystagings/cli/ms/utils/msconfig"
	"github.com/spf13/cobra"
)

func main() {
	config, err := msconfig.LoadConfig()
	if err != nil {
		fmt.Println("Could not load config. Please run:\n\n\t ms configure")
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
		Use:   "ms",
		Short: "Setup your staging environment easily with ms",
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
	var proxyPort int32
	proxyCmd := &cobra.Command{
		Use:   "proxy",
		Short: "proxy to an application inside your staging",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("application name required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.ProxyDeployment(args[0], proxyPort)
		},
	}
	rootCmd.AddCommand(proxyCmd)
	proxyCmd.Flags().Int32VarP(&proxyPort, "port", "p", -1, "Listen on specified port")

	// Exec
	rootCmd.AddCommand(&cobra.Command{
		Use:     "exec",
		Short:   "execute a command into an application container",
		Example: "ms exec [APP_NAME] -- [COMMAND]",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("application name required")
			}
			if cmd.ArgsLenAtDash() <= 0 {
				return errors.New("missing exec command. Please run:\nms exec [APP_NAME] -- [COMMAND]")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			argsLenAtDash := cmd.ArgsLenAtDash()
			execArgs := args[argsLenAtDash:]
			return a.ExecDeployment(args[0], execArgs)
		},
	})

	// Logs
	var logsFollow bool
	var logsTimestamps bool
	var logsLimitBytes int64
	var logsTail int64
	var logsSinceTime string
	var logsSinceSeconds time.Duration
	logsCmd := &cobra.Command{
		Use:     "logs",
		Aliases: []string{"l"},
		Short:   "Print the logs for an appliation",
		Example: "ms logs",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("application name required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.LogsDeployment(args[0], orchestratorcli.LogsOptions{
				Follow:       logsFollow,
				Timestamps:   logsTimestamps,
				LimitBytes:   logsLimitBytes,
				Tail:         logsTail,
				SinceTime:    logsSinceTime,
				SinceSeconds: logsSinceSeconds,
			})
		},
	}
	rootCmd.AddCommand(logsCmd)
	logsCmd.Flags().BoolVarP(&logsFollow, "follow", "f", false, "Specify if the logs should be streamed.")
	logsCmd.Flags().BoolVar(&logsTimestamps, "timestamps", false, "Include timestamps on each line in the log output")
	logsCmd.Flags().Int64Var(&logsLimitBytes, "limit-bytes", 0, "Maximum bytes of logs to return. Defaults to no limit.")
	logsCmd.Flags().Int64Var(&logsTail, "tail", 10, "Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwise 10, if a selector is provided.")
	logsCmd.Flags().StringVar(&logsSinceTime, "since-time", "", "Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used.")
	logsCmd.Flags().DurationVar(&logsSinceSeconds, "since", 0, "Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one of since-time / since may be used.")

	// Status
	rootCmd.AddCommand(&cobra.Command{
		Use:     "status",
		Aliases: []string{"ps"},
		Short:   "get application statuses inside your staging",
		Example: "ms status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return a.Status()
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
