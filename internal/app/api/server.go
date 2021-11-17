package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

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
	logrus.Println("Server start running on port :3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		logrus.Println("Failed to start the server")
		return err
	}
	return nil
}

func (s *Serve) Stop() {
	// do something
}
