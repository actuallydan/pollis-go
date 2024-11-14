package main

import (
	"net/http"

	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
)

// CORSMiddleware wraps an http.HandlerFunc and applies CORS headers
func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	}
}

// AuthMiddleware creates a new middleware handler for authentication
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Use Clerk's WithHeaderAuthorization middleware
		handler := clerkhttp.WithHeaderAuthorization()(next)

		// Call the Clerk-wrapped handler
		handler.ServeHTTP(w, r)
	}
}

func RequiredAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler := clerkhttp.RequireHeaderAuthorization()(next)
		handler.ServeHTTP(w, r)
	}
}

// Middleware chains multiple middleware functions
func Middleware(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		wrappedHandler := handler
		for _, middleware := range middlewares {
			wrappedHandler = middleware(wrappedHandler)
		}

		wrappedHandler.ServeHTTP(w, r)
	}
}
