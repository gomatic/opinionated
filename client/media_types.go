//************************************************************************//
// API "opinionated": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/gomatic/opinionated/design
// --out=$(GOPATH)/src/github.com/gomatic/opinionated
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// Credentials (default view)
//
// Identifier: application/json; view=default
type JSON struct {
	// Username
	ID string `form:"id" json:"id" xml:"id"`
}

// Validate validates the JSON media type instance.
func (mt *JSON) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	return
}

// DecodeJSON decodes the JSON instance encoded in resp body.
func (c *Client) DecodeJSON(resp *http.Response) (*JSON, error) {
	var decoded JSON
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
