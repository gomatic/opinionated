package application

import (
	"fmt"
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
	static := strconv.Itoa(Settings.Port + 1)
	go func() {
		fmt.Println("files on " + static)
		http.ListenAndServe(":"+static, http.FileServer(http.Dir("public/")))
	}()

	port := strconv.Itoa(Settings.Port)
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}

	ctx := context.Background()

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

	//

	fmt.Println("listening on " + port)
	return srv.ListenAndServe()
}
