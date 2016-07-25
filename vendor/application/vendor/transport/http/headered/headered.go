package headered

import (
	"net/http"

	chain "transport/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

//
type Headered func(http.Header)

//
func New(hf Headered) chain.ChainResponseEncoders {
	return func(next httptransport.EncodeResponseFunc) httptransport.EncodeResponseFunc {
		return func(ctx context.Context, resp http.ResponseWriter, i interface{}) error {
			hf(resp.Header())
			return next(ctx, resp, i)
		}
	}
}
