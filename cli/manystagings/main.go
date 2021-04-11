package main

import (
	"log"
	"os"

	"github.com/carlosstrand/manystagings/cli/manystagings/actions"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "configure",
		Usage: "configure the manystagins CLI",
		Commands: []cli.Command{
			{
				Name:    "configure",
				Aliases: []string{"c"},
				Usage:   "configure the manystagins CLI",
				Action: func(c *cli.Context) error {
					return actions.ConfiguteAction()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
