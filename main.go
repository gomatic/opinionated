package main

import (
	"crypto/tls"
	"fmt"
	stderr "log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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

				privPort := strconv.Itoa(Settings.Port)
				pubPort := "80"
				if Settings.Port != 443 || Settings.Port >= 3443 {
					pubPort = strconv.Itoa(Settings.Port - 363)
				}

				// Create service
				srv := service()

				pubAddr := Settings.Addr + ":" + pubPort

				switch Settings.Secure {
				case true:
					crt := filepath.Join(Settings.Program.Data, "server.crt")
					key := filepath.Join(Settings.Program.Data, "server.key")

					_, err := tls.LoadX509KeyPair(crt, key)
					if err == nil {

						{
							routes := http.NewServeMux()
							pub := &http.Server{
								Addr:         pubAddr,
								ReadTimeout:  5 * time.Second,
								WriteTimeout: 10 * time.Second,
								Handler:      routes,
							}

							routes.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
								host := strings.Split(req.Host, ":")
								forward := fmt.Sprintf("https://%s:%s%s", host[0], privPort, req.RequestURI)
								stderr.Println("Forwarding to " + forward)
								http.Redirect(w, req, forward, http.StatusMovedPermanently)
							})
							stderr.Printf("HTTP %s\n", pub.Addr)
							go pub.ListenAndServe()
						}

						privAddr := Settings.Addr + ":" + privPort
						stderr.Printf("HTTPS %s\n", privAddr)
						return srv.ListenAndServeTLS(privAddr, crt, key)
					}
					stderr.Println(err)
					fallthrough

				case false:
					stderr.Printf("HTTP %s\n", pubAddr)
					return srv.ListenAndServe(pubAddr)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
