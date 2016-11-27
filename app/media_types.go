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

package app

import "github.com/goadesign/goa"

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
