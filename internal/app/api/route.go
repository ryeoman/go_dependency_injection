package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type SearchHandlerProvider interface {
	HandleGetSuggestion(w http.ResponseWriter, r *http.Request)
	HandleGetSynonym(w http.ResponseWriter, r *http.Request)
	HandleGetSimilarSpelling(w http.ResponseWriter, r *http.Request)
	HandleGetSimilarSound(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
	SearchHandler SearchHandlerProvider
}

type Route struct {
	// add dependency for router here, like class/struct that handle middleware
	router   RouterProvider
	handlers Handlers
}

type RouterProvider interface {
	chi.Router
}

func NewRouter(router RouterProvider, handlers Handlers) *Route {
	return &Route{
		router:   router,
		handlers: handlers,
	}
}

func (r *Route) RegisterRoute() {
	r.router.Use(middleware.Logger)
	r.router.Get("/suggestion", r.handlers.SearchHandler.HandleGetSuggestion)
}

func (r *Route) GetRouter() RouterProvider {
	return r.router
}
