package main

import (
	"context"

	"go.uber.org/fx"

	restApp "github.com/ryeoman/go_dependency_injection/internal/app/http"
	"github.com/sirupsen/logrus"
)

func register(lifecycle fx.Lifecycle, server *restApp.Serve) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				return server.Start()
			},
			OnStop: func(context.Context) error {
				logrus.Println("app is stopping")
				return nil
			},
		},
	)
}
