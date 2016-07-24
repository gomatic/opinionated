package application

import "net/url"

//
var Settings struct {
	Debugger bool
	Mode     string

	Port int

	Server  string
	Powered string
	Version string

	Docs *url.URL
}
