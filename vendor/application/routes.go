package application

import (
	"middleware/logging"
	"service/login"
	"service/testing"
	"transport/http/caching"
	"transport/http/content"
	"transport/http/server"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"

	"golang.org/x/net/context"
)

//
func routeServices() *mux.Router {
	r := mux.NewRouter()
	ctx := context.Background()

	// Login

	routeLoginService(ctx, r.PathPrefix("/login").Subrouter().StrictSlash(true))

	// Testing

	routeTestService(ctx, r.PathPrefix("/test").Subrouter().StrictSlash(true))

	return r
}

//
func routeLoginService(ctx context.Context, r *mux.Router) {

	logged := logging.New("login").Middleware()
	servered := server.New(Settings.Server+"/auth", Settings.Powered+"/auth")
	json := content.New("application/json")

	loginService := login.New()

	// POST    /

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		ctx,
		logged(loginService.Endpoint()),
		loginService.Decoder,
		servered(json(loginService.Encoder)),
	)).Name("POST")

}

//
func routeTestService(ctx context.Context, r *mux.Router) {

	logged := logging.New("test").Middleware()
	servered := server.New(Settings.Server, Settings.Powered)
	json := content.New("application/json")

	testService := testing.New()

	// POST    /

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		ctx,
		logged(testService.Post()),
		testService.DecodePost,
		servered(json(testService.EncodePost)),
	)).Name("POST")

	// GET     /:id

	r.Methods("POST", "GET").Path("/{id}").Handler(caching.New(-1, httptransport.NewServer(
		ctx,
		logged(testService.Get()),
		testService.DecodeGet,
		servered(json(testService.EncodeGet)),
	))).Name("GET/POST Id")

	// PUT     /:id

	r.Methods("PUT").Path("/{id}").Handler(caching.New(-1, httptransport.NewServer(
		ctx,
		logged(testService.Put()),
		testService.DecodePut,
		servered(json(testService.EncodePut)),
	))).Name("PUT Id")

	// PATCH   /:id

	r.Methods("PATCH").Path("/{id}").Handler(httptransport.NewServer(
		ctx,
		logged(testService.Patch()),
		testService.DecodePatch,
		servered(json(testService.EncodePatch)),
	)).Name("PATCH Id")

	// DELETE  /:id

	r.Methods("DELETE").Path("/{id}").Handler(caching.New(-1, httptransport.NewServer(
		ctx,
		logged(testService.Delete()),
		testService.DecodeDelete,
		servered(json(testService.EncodeDelete)),
	))).Name("DELETE Id")

}
