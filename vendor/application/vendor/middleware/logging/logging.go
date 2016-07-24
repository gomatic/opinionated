package logging

import (
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

	"golang.org/x/net/context"
)

//
type logger struct {
	*log.Context
}

//
func New(name string) *logger {
	l := log.NewLogfmtLogger(os.Stderr)
	c := log.NewContext(l).With("time", log.DefaultTimestampUTC)
	if name != "" {
		c = c.With("service", name)
	}
	return &logger{c}
}

//
func (l *logger) Middleware() endpoint.Middleware {
	if l == nil {
		l = New("")
	}

	return func(next endpoint.Endpoint) endpoint.Endpoint {
		logger := log.Logger(l.Context)
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint", next)
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
