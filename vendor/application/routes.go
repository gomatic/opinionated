package application

import (
	"middleware/logging"
	"net/http"
	"service/login"
	"service/testing"
	"transport/http/caching"
	"transport/http/headered"

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

var (
	servered = headered.New(func(hs http.Header) {
		hs.Set("Server", Settings.Server)
		hs.Set("X-Powered-By", Settings.Powered)
	})
	jsoned = headered.New(func(hs http.Header) {
		hs.Set("Content-Type", "application/json")
	})
)

//
func routeLoginService(ctx context.Context, r *mux.Router) {

	logged := logging.New("login").Middleware()

	loginService := login.New()

	// POST    /

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		ctx,
		logged(loginService.Endpoint()),
		loginService.Decoder,
		servered(jsoned(loginService.Encoder)),
	)).Name("POST")

}

//
func routeTestService(ctx context.Context, r *mux.Router) {

	logged := logging.New("test").Middleware()

	testService := testing.New()

	// POST    /

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		ctx,
		logged(testService.Post()),
		testService.DecodePost,
		servered(jsoned(testService.EncodePost)),
	)).Name("POST")

	// GET     /:id

	r.Methods("POST", "GET").Path("/{id}").Handler(caching.New(-1, httptransport.NewServer(
		ctx,
		logged(testService.Get()),
		testService.DecodeGet,
		servered(jsoned(testService.EncodeGet)),
	))).Name("GET/POST Id")

	// PUT     /:id

	r.Methods("PUT").Path("/{id}").Handler(caching.New(-1, httptransport.NewServer(
		ctx,
		logged(testService.Put()),
		testService.DecodePut,
		servered(jsoned(testService.EncodePut)),
	))).Name("PUT Id")

	// PATCH   /:id

	r.Methods("PATCH").Path("/{id}").Handler(httptransport.NewServer(
		ctx,
		logged(testService.Patch()),
		testService.DecodePatch,
		servered(jsoned(testService.EncodePatch)),
	)).Name("PATCH Id")

	// DELETE  /:id

	r.Methods("DELETE").Path("/{id}").Handler(caching.New(-1, httptransport.NewServer(
		ctx,
		logged(testService.Delete()),
		testService.DecodeDelete,
		servered(jsoned(testService.EncodeDelete)),
	))).Name("DELETE Id")

}
