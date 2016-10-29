package application

import (
	"net/http"
	"transport/http/headered"

	"github.com/gorilla/mux"

	"golang.org/x/net/context"
)

//
func routeServices() *mux.Router {
	r := mux.NewRouter()
	ctx := context.Background()

	// Login

	routeLoginService(ctx, r.PathPrefix("/login").Subrouter().StrictSlash(true))

	// Query

	routeQueryService(ctx, r.PathPrefix("/query").Subrouter().StrictSlash(true))

	// StarWars

	routeStarWarsService(ctx, r.PathPrefix("/starwars").Subrouter().StrictSlash(true))

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
