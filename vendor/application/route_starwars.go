package application

import (
	"middleware/logging"
	"service/starwars"
	"transport/http/caching"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"

	"golang.org/x/net/context"
)

//
func routeStarWarsService(ctx context.Context, r *mux.Router) {

	logger := logging.New("starwars", Settings.Program.Name)
	logged := logger.Middleware()

	starWarsService, err := starwars.New()
	if err != nil {
		logger.Log("error", err.Error())
		return
	}

	// POST     /:query

	r.Methods("POST").Path("/").Handler(caching.New(-1, httptransport.NewServer(
		ctx,
		logged(starWarsService.Endpoint()),
		starWarsService.Decoder,
		servered(jsoned(starWarsService.Encoder)),
	))).Name("POST")

	// GET     /:query

	r.Methods("GET").Path("/{query}").Handler(caching.New(-1, httptransport.NewServer(
		ctx,
		logged(starWarsService.Endpoint()),
		starWarsService.Decoder,
		servered(jsoned(starWarsService.Encoder)),
	))).Name("GET Query")

}
