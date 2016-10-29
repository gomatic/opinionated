package application

import (
	"crypto/tls"
	"fmt"
	stderr "log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"service/view"
)

//
func Start() error {

	// Serve public/

	// Configure endpoints

	routes := routeServices()

	routes.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	routes.PathPrefix("/views/").Handler(http.StripPrefix("/views/", view.Server(http.Dir("views/"))))

	// Start the server

	privPort := strconv.Itoa(Settings.Port)
	pubPort := "80"
	if Settings.Port != 443 || Settings.Port >= 3443 {
		pubPort = strconv.Itoa(Settings.Port - 363)
	}
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      routes,
	}

	if Settings.Insecure {

		srv.Addr = Settings.Addr + ":" + pubPort
		stderr.Printf("HTTP %s\n", srv.Addr)
		return srv.ListenAndServe()

	} else {
		crt := filepath.Join(Settings.Program.Data, "server.crt")
		key := filepath.Join(Settings.Program.Data, "server.key")

		if cert, err := tls.LoadX509KeyPair(crt, key); err != nil {
			stderr.Println(err)

			srv.Addr = Settings.Addr + ":" + pubPort
			stderr.Printf("HTTP %s\n", srv.Addr)
			return srv.ListenAndServe()

		} else {

			srv.TLSConfig = &tls.Config{
				Certificates: []tls.Certificate{
					cert,
				},
			}

			{
				routes := http.NewServeMux()
				pub := &http.Server{
					Addr:         Settings.Addr + ":" + pubPort,
					ReadTimeout:  srv.ReadTimeout,
					WriteTimeout: srv.WriteTimeout,
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

			srv.Addr = Settings.Addr + ":" + privPort
			stderr.Printf("HTTPS %s\n", srv.Addr)
			return srv.ListenAndServeTLS(crt, key)
		}
	}
}
