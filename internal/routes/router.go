package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct {
	mux *chi.Mux
}

func NewRouter() *Router {
	mux := chi.NewRouter()

	r := &Router{
		mux: mux,
	}

	return r
}

func NewDefaultRouter() *Router {
	r := NewRouter()

	// Middleware which will not change the response.
	r.WithMiddleware(middleware.RequestID)

	// Logger must come before any other middleware that may change
	// the response, such as `middleware.Recoverer`.
	r.WithMiddleware(middleware.Logger)

	// Middleware which may change the response.
	r.WithMiddleware(middleware.Recoverer)

	return r
}

func (r *Router) GetMux() *chi.Mux {
	return r.mux
}

func (r *Router) WithMiddleware(h func(http.Handler) http.Handler) {
	r.mux.Use(h)
}

func (r *Router) WithGet(p string, f http.HandlerFunc) {
	r.mux.Get(p, f)
}
