package view

import "net/http"

//
type FileSystem interface {
}

//
type fileHandler struct {
	root FileSystem
}

//
func Server(root FileSystem) http.Handler {
	return &fileHandler{root}
}

//
func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
