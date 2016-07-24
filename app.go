package main

import (
	_ "expvar"
	"log"
	"os"

	"application"

	"github.com/urfave/cli"
)

//
const MAJOR = "0.1"

// DO NOT UPDATE. This is populated by the build. See the Makefile.
var VERSION = "0"

//
func main() {
	app := cli.NewApp()
	app.Name = "opinionated"
	app.Usage = "Opinionated."
	app.Version = MAJOR + "." + VERSION
	app.EnableBashCompletion = true

	application.Settings.Version = app.Version
	application.Settings.Server = app.Name
	application.Settings.Powered = app.Name + "/" + app.Version

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, debugger, debugging, d",
			Usage:       "Enable debugging server.",
			Destination: &application.Settings.Debugger,
		},
	}

	app.Before = func(ctx *cli.Context) error {
		if application.Settings.Debugger {
			application.Debugger()
		}
		return nil
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:    "configuration",
			Aliases: []string{"configure", "config"},
			Usage:   "Configure.",
			Action: func(ctx *cli.Context) error {
				log.Println(application.Configure())
				return nil
			},
		},
		cli.Command{
			Name:    "manager",
			Aliases: []string{"manage", "administer", "admin"},
			Usage:   "Administer.",
			Action: func(ctx *cli.Context) error {
				log.Println(application.Manage())
				return nil
			},
		},
		cli.Command{
			Name:    "start",
			Aliases: []string{"server", "serve", "run"},
			Usage:   "Run.",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:        "port, p",
					Usage:       "Server port.",
					Value:       3000,
					Destination: &application.Settings.Port,
				},
			},
			Action: func(ctx *cli.Context) error {
				log.Println(application.Start())
				return nil
			},
		},
	}

	app.Run(os.Args)
}
