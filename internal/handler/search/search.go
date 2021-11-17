package search

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ryeoman/go_dependency_injection/internal/handler"
	"github.com/ryeoman/go_dependency_injection/internal/usecase/autocomplete"
)

type AutoCompleteUsecaseProvider interface {
	GetSuggestion(ctx context.Context, keyword string) (autocomplete.Suggestions, error)
}

type SimilarityUsecaseProvider interface {
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
	suggestions, _ := s.autoCompleteUsecase.GetSuggestion(r.Context(), "misal")
	json.NewEncoder(w).Encode(handler.Response{
		Status: http.StatusOK,
		Data:   suggestions,
	})
}

func (s *Handler) HandleGetSynonym(w http.ResponseWriter, r *http.Request) {

}

func (s *Handler) HandleGetSimilarSpelling(w http.ResponseWriter, r *http.Request) {

}

func (s *Handler) HandleGetSimilarSound(w http.ResponseWriter, r *http.Request) {

}
