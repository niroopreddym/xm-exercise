package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

//ContextMiddleware handles context across the layers
func (handler *AuthHandler) ContextMiddleware(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// FIXME Do something with our context
		timeoutCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
		defer cancel()

		next.ServeHTTP(w, r)

		select {
		case <-timeoutCtx.Done():
			fmt.Println("context done exceeded the deadline:", ctx.Err())
		}
	})
}
