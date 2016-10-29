package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/graphql-go/graphql"
	"golang.org/x/net/context"
)

// Errors

// ErrNoQuery.
var ErrNoQuery = errors.New("empty query")

// Model

// Query operations.
type Service interface {
	Query(string) (*graphql.Result, error)
}

// Private Query model.
type queryService struct{}

// Query requests.
type queryRequest struct {
	S string `json:"query"`
}

var (
	schema graphql.Schema
)

//
func (queryService) Query(query string) (*graphql.Result, error) {
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		return nil, fmt.Errorf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	return r, nil
}

// Initialization

//
type queryEndpoint endpoint.Endpoint

//
func New() (endpoint queryEndpoint, err error) {
	svc := queryService{}

	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "Query", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		return nil, err
	}

	endpoint = func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(queryRequest)
		return svc.Query(req.S)
	}

	return endpoint, nil
}

// Convenience to get a go-kit type back of query's private endpoint type.
func (e queryEndpoint) Endpoint() endpoint.Endpoint {
	return endpoint.Endpoint(e)
}

// Request/response encoding and decoding.

//
func (e queryEndpoint) Decoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request queryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

//
func (e queryEndpoint) Encoder(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
