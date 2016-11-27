//************************************************************************//
// API "opinionated": Application Contexts
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
)

// LoginUserContext provides the user login action context.
type LoginUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Username string
}

// NewLoginUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the user controller login action.
func NewLoginUserContext(ctx context.Context, service *goa.Service) (*LoginUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := LoginUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramUsername := req.Params["username"]
	if len(paramUsername) > 0 {
		rawUsername := paramUsername[0]
		rctx.Username = rawUsername
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *LoginUserContext) OK(r *JSON) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *LoginUserContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
