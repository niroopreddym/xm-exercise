package handlers

import (
	"net/http"

	"github.com/niroopreddym/xm-exercise/helpers"
)

//IPFIlterMiddleware ...
func (handler *AuthHandler) IPFIlterMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// get the real IP of the user, see below
		addr := helpers.GetClientIP(req)

		// the actual vaildation - replace with whatever you want
		if !helpers.IsValidRequest(addr) {
			http.Error(w, "Blocked", 401)
			return
		}

		// pass the request to the mux
		h.ServeHTTP(w, req)
	})
}
