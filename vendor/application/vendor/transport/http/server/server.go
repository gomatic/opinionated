package server

import (
	"net/http"

	chain "transport/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func New(server, powered string) chain.ChainResponseEncoders {
	return func(next httptransport.EncodeResponseFunc) httptransport.EncodeResponseFunc {
		return func(ctx context.Context, resp http.ResponseWriter, i interface{}) error {

			hs := resp.Header()
			hs.Set("Server", server)
			hs.Set("X-Powered-By", powered)

			return next(ctx, resp, i)
		}
	}
}
