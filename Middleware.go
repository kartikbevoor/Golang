package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// Middleware is a function that sits between incoming http request and the final handler
// It can: modify request, responses, stop request, add extra data, perform checks,
// execute logic before/after handler

// Middleware Flow/Request-Response Lifecycle
// Request → Middleware 1 → Middleware 2 → Middleware 3 → Handler → Response
// When response comes back:
// Handler → Middleware 3 → Middleware 2 → Middleware 1 → Client

func Middleware() {

}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// BEFORE handler
		fmt.Println("Before request")

		next.ServeHTTP(w, r)

		// AFTER handler
		fmt.Println("After request")
	})
}

// next http.Handler -> next is the handler that should run after the middleware finishes its pre-processing.
// next http.Handler: the next middleware OR the final handler
// this above middleware returns http handler function
// ServeHTTP(): ServeHTTP(ResponseWriter, *Request) -> it passes control to the next layer.

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before")

		next.ServeHTTP(w, r)

		fmt.Println("After")
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func callHello() {
	handler := http.HandlerFunc(hello)

	http.Handle("/", loggingMiddleware(handler))

	http.ListenAndServe(":8080", nil)
}

// Logging Midlleware
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Method:", r.Method)
		fmt.Println("URL:", r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

// Authentication Middleware
func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Panic recovery middleware: Protects server from crashing.
func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", 500)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// Request Timing Middleware
func timer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		fmt.Println("time:", time.Since(start))
	})
}

// Context Middleware: Middleware often injects data into request context.
// Example:user ID, organization ID, request ID, roles, permissions

type key string

const UserIDKey key = "user_id"

func authId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), UserIDKey, "1001")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAuthId(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(UserIDKey)

	fmt.Println(id)
}

// Middleware in Chi Router
func ChiMiddleware() {
	r := chi.NewRouter()

	r.Use(logging)
	r.Use(auth)
	r.Use(recovery) // Note: order of these middlewares matter
	// r.Handle("/", logging(auth(recovery(http.HandlerFunc(hello),),),))

	r.Get("/", hello)
}

// Route specific middleware
func routeSpecificMiddleware() {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(auth)

		// r.Get("/profile", profileHandler)
		// r.Get("/settings", settingsHandler)
	})
}

// Global Middleware: Applies to all routes
// r.Use(logging)

// NOTE: Middleware order matters, change in order changes behaviour.
// ex: if authentication middleware runs before logging, requests may never be processed ->
// unauthorised process can never be logged

// JWT middleware
func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "missing token in header", 401)
			return
		}

		// validate token here

		next.ServeHTTP(w, r)
	})
}

// Middleware vs Handleer
// | Middleware                    | Handler                  |
// | ----------------------------- | ------------------------ |
// | Executes before/after handler | Final request processor  |
// | Can modify request            | Returns final response   |
// | Can stop request              | Generates business logic |
// | Cross-cutting concern         | Feature-specific logic   |

// Best Practices:
// Keep Middleware Small, Avoid Business Logic, Use Context Carefully, Never Panic Inside Middleware, Always Call next.ServeHTTP()

// Important Middleware patterns
// RequestId middleware
// Rate limiting middleware
// Cors middleware: controls frontend access
// Compression middleware: brotli, gzip
// Metrics Middleware: tracks request count, latency, error

// Middleware in Production Architecture
// Request
//  ↓
// Load Balancer
//  ↓
// Reverse Proxy
//  ↓
// Go Server
//  ↓
// Request ID Middleware
//  ↓
// Logger Middleware
//  ↓
// Recovery Middleware
//  ↓
// JWT Middleware
//  ↓
// Rate Limiter
//  ↓
// Validation Middleware
//  ↓
// Handler
//  ↓
// Service Layer
//  ↓
// Repository Layer
//  ↓
// Database

// Important Interview Questions:
// Why middleware? -> To avoid duplicate logic across handlers.
// Why next.ServeHTTP()? -> To pass control to next middleware/handler.
// Can middleware stop requests? -> Yes, using return statement before calling next.
// Why context in middleware? -> To share request-scoped data safely.
// Difference between middleware and handler? -> Middleware wraps handlers and adds reusable behavior.
