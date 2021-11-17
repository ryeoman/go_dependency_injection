package similarity

import (
	"context"

	"github.com/ryeoman/go_dependency_injection/internal/repository/search"
)

type SearchProvider interface {
	Synonym(ctx context.Context, word string) (search.Words, error)
	SoundsLike(ctx context.Context, word string) (search.Words, error)
	SpelledLike(ctx context.Context, word string) (search.Words, error)
}

type Usecase struct {
	provider SearchProvider
}

func New(provider SearchProvider) *Usecase {
	return &Usecase{
		provider: provider,
	}
}

func (u *Usecase) GetSynonym(ctx context.Context, word string) (Similars, error) {
	similars := Similars{}
	words, err := u.provider.Synonym(ctx, word)
	if err != nil {
		return similars, err
	}
	similars.fromWords(words)
	return similars, nil
}

func (u *Usecase) GetSimilarSpelling(ctx context.Context, word string) (Similars, error) {
	similars := Similars{}
	words, err := u.provider.SpelledLike(ctx, word)
	if err != nil {
		return similars, err
	}
	similars.fromWords(words)
	return similars, nil
}

func (u *Usecase) GetSimilarSound(ctx context.Context, word string) (Similars, error) {
	similars := Similars{}
	words, err := u.provider.SoundsLike(ctx, word)
	if err != nil {
		return similars, err
	}
	similars.fromWords(words)
	return similars, nil
}
