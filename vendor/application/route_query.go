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

	logger := logging.New("query", Settings.Program.Name)
	logged := logger.Middleware()

	queryService, err := query.New()
	if err != nil {
		logger.Log("error", err.Error())
		return
	}

	// POST     /:id

	r.Methods("POST", "GET").Path("/{id}").Handler(caching.New(-1, httptransport.NewServer(
		ctx,
		logged(queryService.Endpoint()),
		queryService.Decoder,
		servered(jsoned(queryService.Encoder)),
	))).Name("GET/POST Id")

}
