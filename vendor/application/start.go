package application

import (
	"crypto/tls"
	"fmt"
	stderr "log"
	"net/http"
	"strconv"
	"time"

	"middleware/logging"
	"service/login"
	"service/testing"
	"transport/http/caching"
	"transport/http/content"
	"transport/http/server"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

//
func Start() error {

	// Serve public/
	go func() {
		static := strconv.Itoa(Settings.Port + 1)
		fmt.Println("files on " + static)
		http.ListenAndServe(":"+static, http.FileServer(http.Dir("public/")))
	}()

	mux := http.NewServeMux()
	ctx := context.Background()

	// Configure endpoints

	// Login

	loginService := login.New()
	loggedLogin := logging.New(nil)(loginService.Endpoint())

	loginEncoder := server.New("opinionated", "gomatic/opinionated")(loginService.Encoder)
	loginEncoder = content.New("application/json")(loginEncoder)

	loginHandler := httptransport.NewServer(
		ctx,
		loggedLogin,
		loginService.Decoder,
		loginEncoder,
	)

	mux.Handle("/login", loginHandler)

	// Testing

	testService := testing.New()
	loggedTest := logging.New(nil)(testService.Endpoint())

	testEncoder := server.New("opinionated", "gomatic/opinionated")(testService.Encoder)
	testEncoder = content.New("application/json")(testEncoder)

	testHandler := httptransport.NewServer(
		ctx,
		loggedTest,
		testService.Decoder,
		testEncoder,
	)

	mux.Handle("/", caching.New(testHandler))

	// Start the server

	port := strconv.Itoa(Settings.Port)
	srv := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}

	if cer, err := tls.LoadX509KeyPair("server.crt", "server.key"); err != nil {
		stderr.Println(err)

	} else {

		srv.TLSConfig = &tls.Config{Certificates: []tls.Certificate{cer}}

	}

	fmt.Println("listening on " + port)
	return srv.ListenAndServe()
}
