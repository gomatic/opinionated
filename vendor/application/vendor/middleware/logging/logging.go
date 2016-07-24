package logging

import (
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"golang.org/x/net/context"
)

//
func New(logger log.Logger) endpoint.Middleware {
	if logger == nil {
		logger = log.NewLogfmtLogger(os.Stderr)
	}
	logger = log.NewContext(logger).With("time", log.DefaultTimestampUTC)

	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint", next)
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
