package main

import (
	"context"
	"fmt"

	"github.com/go-chi/chi"
	restApp "github.com/ryeoman/go_dependency_injection/internal/app/api"
	searchHandler "github.com/ryeoman/go_dependency_injection/internal/handler/search"
	searchRepo "github.com/ryeoman/go_dependency_injection/internal/repository/search"
	"github.com/ryeoman/go_dependency_injection/internal/usecase/autocomplete"
	"github.com/ryeoman/go_dependency_injection/internal/usecase/similarity"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Logger(logrus.New()),
		fx.Provide(
			// repo provider
			provideAutocompleteSearchRepo,
			provideSimilaritySearchRepo,
			// Usecase provider
			provideAutoCompleteUsecase,
			provideSimilarityUsecase,
			// handler provider
			provideHandlers,
			provideSearchHandler,
			// app provider
			provideRoute,
			provideRouter,
			// rovide server
			provideServer,
		),
		fx.Invoke(
			register,
		),
	)

	if err := app.Start(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func provideAutocompleteSearchRepo() autocomplete.SearchProvider {
	return searchRepo.New()
}

func provideSimilaritySearchRepo() similarity.SearchProvider {
	return searchRepo.New()
}

func provideAutoCompleteUsecase(provider autocomplete.SearchProvider) searchHandler.AutoCompleteUsecaseProvider {
	return autocomplete.New(provider)
}

func provideSimilarityUsecase(provider similarity.SearchProvider) searchHandler.SimilarityUsecaseProvider {
	return similarity.New(provider)
}

func provideSearchHandler(
	autoCompleteUsecase searchHandler.AutoCompleteUsecaseProvider,
	similarityUsecase searchHandler.SimilarityUsecaseProvider,
) restApp.SearchHandlerProvider {
	return searchHandler.New(autoCompleteUsecase, similarityUsecase)
}

func provideHandlers(
	searchHandler restApp.SearchHandlerProvider,
) restApp.Handlers {
	return restApp.Handlers{
		SearchHandler: searchHandler,
	}
}

func provideRouter() restApp.RouterProvider {
	return chi.NewRouter()
}

func provideRoute(
	router restApp.RouterProvider,
	handlers restApp.Handlers,
) restApp.RouteProvider {
	return restApp.NewRouter(router, handlers)
}

func provideServer(route restApp.RouteProvider) *restApp.Serve {
	return restApp.NewServer(route)
}
