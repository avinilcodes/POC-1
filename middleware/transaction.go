package middleware

import (
	"context"
	"net/http"
)

func TransactionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), "userID", "b36f11d2-2f8b-46cb-8b45-4b525bba218d")
		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
