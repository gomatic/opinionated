package application

import (
	"middleware/logging"
	"service/testing"
	"transport/http/caching"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"

	"golang.org/x/net/context"
)

//
func routeTestService(ctx context.Context, r *mux.Router) {

	logged := logging.New("test", Settings.Program.Name).Middleware()

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
