package autocomplete

import (
	"context"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/ryeoman/go_dependency_injection/internal/repository/search"
)

//go:generate mockgen -source=./autocomplete.go -destination=./autocomplete_mock.go -package=autocomplete
type SearchProvider interface {
	Suggestion(ctx context.Context, keyword string) (search.Words, error)
}

type Usecase struct {
	provider SearchProvider
}

func New(provider SearchProvider) *Usecase {
	return &Usecase{
		provider: provider,
	}
}

func (u *Usecase) GetSuggestion(ctx context.Context, keyword string) (Suggestions, error) {
	suggestions := Suggestions{}

	// will return when minimum number of char was not fulfilled
	if !u.isValid(keyword) {
		return suggestions, nil
	}

	words, err := u.provider.Suggestion(ctx, keyword)
	if err != nil {
		return suggestions, err
	}

	// transform word object to suggestion object
	suggestions.fromWords(words)

	return suggestions, nil
}

func (u *Usecase) isValid(keyword string) bool {
	return len(keyword) > 2
}
