package application

import (
	"middleware/logging"
	"service/query"
	"transport/http/caching"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"

	"golang.org/x/net/context"
)

//
func routeQueryService(ctx context.Context, r *mux.Router) {

	logged := logging.New("query", Settings.Program.Name).Middleware()

	queryService := query.New()

	// GET     /:id

	r.Methods("GET").Path("/{id}").Handler(caching.New(-1, httptransport.NewServer(
		ctx,
		logged(queryService.Endpoint()),
		queryService.Decoder,
		servered(jsoned(queryService.Encoder)),
	))).Name("GET Id")

}
