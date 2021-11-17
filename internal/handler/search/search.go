package search

import (
	"context"
	"net/http"

	"github.com/ryeoman/go_dependency_injection/internal/handler"
	"github.com/ryeoman/go_dependency_injection/internal/usecase/autocomplete"
	"github.com/ryeoman/go_dependency_injection/internal/usecase/similarity"
)

type AutoCompleteUsecaseProvider interface {
	GetSuggestion(ctx context.Context, keyword string) (autocomplete.Suggestions, error)
}

type SimilarityUsecaseProvider interface {
	GetSynonym(ctx context.Context, word string) (similarity.Similars, error)
	GetSimilarSpelling(ctx context.Context, word string) (similarity.Similars, error)
	GetSimilarSound(ctx context.Context, word string) (similarity.Similars, error)
}

type Handler struct {
	autoCompleteUsecase AutoCompleteUsecaseProvider
	similarityUsecase   SimilarityUsecaseProvider
}

func New(
	autoCompleteUsecase AutoCompleteUsecaseProvider,
	similarityUsecase SimilarityUsecaseProvider,
) *Handler {
	return &Handler{
		autoCompleteUsecase: autoCompleteUsecase,
		similarityUsecase:   similarityUsecase,
	}
}

func (s *Handler) HandleGetSuggestion(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = r.Context()
		query = r.URL.Query()
		word  = query.Get("word")
	)

	suggestions, err := s.autoCompleteUsecase.GetSuggestion(r.Context(), word)
	if err != nil {
		handler.WriteJSONResponse(ctx, w, nil, err, http.StatusInternalServerError)
		return
	}

	handler.WriteJSONResponse(ctx, w, suggestions, nil, http.StatusOK)
}

func (s *Handler) HandleGetSynonym(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = r.Context()
		query = r.URL.Query()
		word  = query.Get("word")
	)

	suggestions, err := s.similarityUsecase.GetSynonym(r.Context(), word)
	if err != nil {
		handler.WriteJSONResponse(ctx, w, nil, err, http.StatusInternalServerError)
		return
	}

	handler.WriteJSONResponse(ctx, w, suggestions, nil, http.StatusOK)
}

func (s *Handler) HandleGetSimilarSpelling(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = r.Context()
		query = r.URL.Query()
		word  = query.Get("word")
	)

	suggestions, err := s.similarityUsecase.GetSimilarSpelling(r.Context(), word)
	if err != nil {
		handler.WriteJSONResponse(ctx, w, nil, err, http.StatusInternalServerError)
		return
	}

	handler.WriteJSONResponse(ctx, w, suggestions, nil, http.StatusOK)
}

func (s *Handler) HandleGetSimilarSound(w http.ResponseWriter, r *http.Request) {
	var (
		ctx   = r.Context()
		query = r.URL.Query()
		word  = query.Get("word")
	)

	suggestions, err := s.similarityUsecase.GetSimilarSound(r.Context(), word)
	if err != nil {
		handler.WriteJSONResponse(ctx, w, nil, err, http.StatusInternalServerError)
		return
	}

	handler.WriteJSONResponse(ctx, w, suggestions, nil, http.StatusOK)
}
