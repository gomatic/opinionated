//go:generate goagen bootstrap -d github.com/gomatic/opinionated/design

package main

import (
	"crypto/tls"
	stderr "log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	op "github.com/gomatic/opinionated/app"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "opinionated"
	app.Usage = "Opinionated."
	app.Version = MAJOR + "." + VERSION
	app.EnableBashCompletion = true

	Settings.Version = app.Version
	Settings.Server = app.Name
	Settings.Powered = app.Name + "/" + app.Version

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, debugger, debugging, d",
			Usage:       "Enable debugging server.",
			Destination: &Settings.Debugger,
		},
	}

	app.Before = func(ctx *cli.Context) error {
		if Settings.Debugger {
			Debugger()
		}
		Settings.Mode = strings.ToLower(Settings.Mode)
		if Settings.Addr == "" && Settings.Mode == "development" {
			Settings.Addr = "localhost"
		}
		return nil
	}

	app.Commands = []cli.Command{
		cli.Command{
			Name:    "start",
			Aliases: []string{"server", "serve", "run"},
			Usage:   "Run.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "address, a",
					Usage:       "Bind to address.",
					Value:       "localhost",
					Destination: &Settings.Addr,
				},
				cli.IntFlag{
					Name:        "port, p",
					Usage:       "Server port.",
					Value:       3443,
					Destination: &Settings.Port,
				},
				cli.BoolFlag{
					Name:        "secure, k",
					Usage:       "Enable TLS server.",
					Destination: &Settings.Secure,
				},
				cli.BoolFlag{
					Name:        "verbose, v",
					Usage:       "Enable verbose server.",
					Destination: &Settings.Verbose,
				},
				cli.StringFlag{
					Name:        "mode",
					Usage:       "Provide a run-mode.",
					EnvVar:      "DEPLOYMENT",
					Destination: &Settings.Mode,
				},
			},
			Before: func(ctx *cli.Context) error {
				Settings.Mode = strings.ToLower(Settings.Mode)
				if Settings.Addr == "" && Settings.Mode == "development" {
					Settings.Addr = "localhost"
				}
				return nil
			},
			Action: func(ctx *cli.Context) error {

				// Create service
				service := goa.New("opinionated")

				// Mount middleware
				service.Use(middleware.RequestID())
				service.Use(middleware.LogRequest(true))
				service.Use(middleware.ErrorHandler(service, true))
				service.Use(middleware.Recover())

				// Mount "user" controller
				c := NewUserController(service)
				op.MountUserController(service, c)

				privPort := strconv.Itoa(Settings.Port)
				pubPort := "80"
				if Settings.Port != 443 || Settings.Port >= 3443 {
					pubPort = strconv.Itoa(Settings.Port - 363)
				}

				switch Settings.Secure {
				case true:
					crt := filepath.Join(Settings.Program.Data, "server.crt")
					key := filepath.Join(Settings.Program.Data, "server.key")

					_, err := tls.LoadX509KeyPair(crt, key)
					if err == nil {
						addr := Settings.Addr + ":" + privPort
						stderr.Printf("HTTPS %s\n", addr)
						return service.ListenAndServeTLS(addr, crt, key)
					}
					stderr.Println(err)
					fallthrough

				case false:
					addr := Settings.Addr + ":" + pubPort
					stderr.Printf("HTTP %s\n", addr)
					return service.ListenAndServe(addr)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
