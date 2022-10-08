package handlers

import "net/http"

//MiddlewareIface abstracts the functionlities on all middleware handler
type MiddlewareIface interface {
	GenerateToken(w http.ResponseWriter, r *http.Request)
}
