//************************************************************************//
// API "opinionated": Application Controllers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/gomatic/opinionated/design
// --out=$(GOPATH)/src/github.com/gomatic/opinionated
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// UserController is the controller interface for the User actions.
type UserController interface {
	goa.Muxer
	Login(*LoginUserContext) error
}

// MountUserController "mounts" a User resource controller on the given service.
func MountUserController(service *goa.Service, ctrl UserController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewLoginUserContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Login(rctx)
	}
	service.Mux.Handle("GET", "/u/:username", ctrl.MuxHandler("Login", h, nil))
	service.LogInfo("mount", "ctrl", "User", "action", "Login", "route", "GET /u/:username")
}
