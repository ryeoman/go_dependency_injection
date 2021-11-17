package main

import (
	"context"
	"log"

	"go.uber.org/fx"

	restApp "github.com/ryeoman/go_dependency_injection/internal/app/api"
)

func register(lifecycle fx.Lifecycle, server *restApp.Serve) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				return server.Start()
			},
			OnStop: func(context.Context) error {
				log.Println("app stoping")
				return nil
			},
		},
	)
}
