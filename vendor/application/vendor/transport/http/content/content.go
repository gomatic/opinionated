package content

import (
	"net/http"

	chain "transport/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func New(contentType string) chain.ChainResponseEncoders {
	return func(next httptransport.EncodeResponseFunc) httptransport.EncodeResponseFunc {
		return func(ctx context.Context, resp http.ResponseWriter, i interface{}) error {

			hs := resp.Header()
			hs.Set("Content-Type", contentType)

			return next(ctx, resp, i)
		}
	}
}
