package application

import (
	"crypto/tls"
	"fmt"
	stderr "log"
	"net/http"
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

	fmt.Println("listening on " + srv.Addr)

	if Settings.Insecure {

		return srv.ListenAndServe()

	} else if cert, err := tls.LoadX509KeyPair("server.crt", "server.key"); err != nil {
		stderr.Println(err)

		return srv.ListenAndServe()

	} else {

		srv.TLSConfig = &tls.Config{
			Certificates: []tls.Certificate{
				cert,
			},
		}

		return srv.ListenAndServeTLS("server.crt", "server.key")
	}

}
