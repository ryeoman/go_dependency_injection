package api

import "net/http"

type RouteProvider interface {
	RegisterRoute()
	GetRouter() RouterProvider
}

type Serve struct {
	// add dependency for server here
	router RouteProvider
}

func NewServer(router RouteProvider) *Serve {
	return &Serve{
		router: router,
	}
}

func (s *Serve) Start() error {
	s.router.RegisterRoute()
	router := s.router.GetRouter()
	return http.ListenAndServe(":3000", router)
}

func (s *Serve) Stop() {
	// do something
}
