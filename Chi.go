package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Chi is a lightweight, fast, and idiomatic HTTP router used to build web APIs and backend services.
// it stays very close to Go’s standard library while still providing powerful routing and middleware support.

// Definition:
// Chi is a router and middleware framework built on top of Go's standard library net/http.
// It helps in:
// Route HTTP requests, Organize APIs, Use middleware, Build REST services, Create scalable backend systems

func Chi() {

}

func basicChiRouter() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) { // (route, handler function)
		fmt.Fprintf(w, "Hello")
	})

	http.ListenAndServe(":8080", r) // ListenAndServe(address, handler)
}

// Using URL parameters
func getUser(w http.ResponseWriter, r *http.Request) { // handler function

	id := chi.URLParam(r, "id")
	id2 := r.URL.Query().Get("id") // here r is request

	fmt.Fprintf(w, "User ID: %s, %s", id, id2)
}

// Route Groups
func RouteGroups() {
	r := chi.NewRouter()

	r.Route("/users", func(r chi.Router) { // .Route(): creates a sub-router whose base path is

		r.Get("/", getUsers)

		r.Get("/{id}", getUser)

		r.Get("{id}/transaction", getTransactions)

	})

	r.Mount("/users", userRoutes()) // subrouter
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Transanctions")
}

// Nested Routes
func NestedRoutes() {

	// r.Route("/accounts", func(r chi.Router) {

	// 	r.Get("/", listAccounts)

	// 	r.Route("/{accountID}", func(r chi.Router) {

	// 		r.Get("/", getAccount)
	// 		r.Get("/transactions", getTransactions)

	// 	})

	// })
}

// Middleware in Chi:
// Middleware runs before or after a request handler.
// uses: Authentication, Logging, Rate limiting, Request ID, Recovery from panic

// Subrouters: Subrouters allow modular API design.
func userRoutes() http.Handler {

	r := chi.NewRouter()

	r.Get("/", getUsers)
	r.Get("/{id}", getUser)

	return r
}

// Context in Chi: Chi uses Go's context system.
// Context is used for: request lifecycle, deadlines, database queries, authentication data

func handlerContext(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	fmt.Println(ctx)
	fmt.Fprintf(w, "ctx")
}
