package http

import "net/http"

// JSONMiddleware is a middleware that adds the JSON content type to the response header,
// before passing the request to the next handler in the chain.
// This middleware ensures that the response is interpreted as JSON by the client.
func JSONMiddleware(next http.Handler) http.Handler {
	// Return a http.HandlerFunc as the http.Handler.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the response header to have content type as JSON
		w.Header().Set("Contet-Type", "application/json; charset=UTF-8")
		// Call the ServeHTTP method of the next http.Handler in the chain
		next.ServeHTTP(w, r)
	})
}
