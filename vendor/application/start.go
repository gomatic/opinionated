package application

import (
	"crypto/tls"
	"fmt"
	stderr "log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

//
func Start() error {

	// Serve public/

	// Configure endpoints

	mux := routeServices()

	mux.Handle("/", http.FileServer(http.Dir("public/")))

	// Start the server

	srv := &http.Server{
		Addr:         Settings.Addr + ":" + strconv.Itoa(Settings.Port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}

	message := "listening on " + srv.Addr

	if Settings.Insecure {

		fmt.Printf("HTTP %s\n", message)
		return srv.ListenAndServe()

	} else {
		crt := filepath.Join(Settings.Program.Data, "server.crt")
		key := filepath.Join(Settings.Program.Data, "server.key")

		if cert, err := tls.LoadX509KeyPair(crt, key); err != nil {
			stderr.Println(err)

			fmt.Printf("HTTP %s\n", message)
			return srv.ListenAndServe()

		} else {

			srv.TLSConfig = &tls.Config{
				Certificates: []tls.Certificate{
					cert,
				},
			}

			fmt.Printf("HTTPS %s\n", message)
			return srv.ListenAndServeTLS(crt, key)
		}
	}
}
