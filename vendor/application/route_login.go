package application

import (
	"middleware/logging"
	"service/login"

	"github.com/gorilla/mux"

	httptransport "github.com/go-kit/kit/transport/http"

	"golang.org/x/net/context"
)

//
func routeLoginService(ctx context.Context, r *mux.Router) {

	logged := logging.New("login", Settings.Program.Name).Middleware()

	loginService := login.New()

	// POST    /

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		ctx,
		logged(loginService.Endpoint()),
		loginService.Decoder,
		servered(jsoned(loginService.Encoder)),
	)).Name("POST")

}
