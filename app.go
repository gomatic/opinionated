package main

import (
	_ "expvar"
	"os"
	"strings"

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
		application.Settings.Mode = strings.ToLower(application.Settings.Mode)
		if application.Settings.Addr == "" && application.Settings.Mode == "development" {
			application.Settings.Addr = "localhost"
		}
		return nil
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:    "configuration",
			Aliases: []string{"configure", "config"},
			Usage:   "Configure.",
			Action: func(ctx *cli.Context) error {
				return application.Configure()
			},
		},
		cli.Command{
			Name:    "manager",
			Aliases: []string{"manage", "administer", "admin"},
			Usage:   "Administer.",
			Action: func(ctx *cli.Context) error {
				return application.Manage()
			},
		},
		cli.Command{
			Name:    "start",
			Aliases: []string{"server", "serve", "run"},
			Usage:   "Run.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "address, a",
					Usage:       "Bind to address.",
					Value:       "",
					Destination: &application.Settings.Addr,
				},
				cli.IntFlag{
					Name:        "port, p",
					Usage:       "Server port.",
					Value:       3443,
					Destination: &application.Settings.Port,
				},
				cli.BoolFlag{
					Name:        "insecure, k",
					Usage:       "Disable TLS server.",
					Destination: &application.Settings.Insecure,
				},
				cli.BoolFlag{
					Name:        "verbose, v",
					Usage:       "Enable verbose server.",
					Destination: &application.Settings.Verbose,
				},
				cli.StringFlag{
					Name:        "mode",
					Usage:       "Provide a run-mode.",
					EnvVar:      "DEPLOYMENT",
					Destination: &application.Settings.Mode,
				},
			},
			Before: func(ctx *cli.Context) error {
				application.Settings.Mode = strings.ToLower(application.Settings.Mode)
				if application.Settings.Addr == "" && application.Settings.Mode == "development" {
					application.Settings.Addr = "localhost"
				}
				return nil
			},
			Action: func(ctx *cli.Context) error {
				return application.Start()
			},
		},
	}

	app.Run(os.Args)
}
