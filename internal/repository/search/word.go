package search

import (
	"context"

	"github.com/kostaspt/go-datamuse/v2"
)

type Search struct {
}

func New() *Search {
	return &Search{}
}

func (s *Search) SoundsLike(ctx context.Context, word string) (Words, error) {
	words := Words{}
	results, err := datamuse.New().Words().SoundsLike(word).Get()
	if err != nil {
		return words, err
	}
	words.fromResults(results)
	return words, nil
}

func (s *Search) SpelledLike(ctx context.Context, word string) (Words, error) {
	words := Words{}
	results, err := datamuse.New().Words().SpelledLike(word).Get()
	if err != nil {
		return words, err
	}
	words.fromResults(results)
	return words, nil
}

func (s *Search) Synonym(ctx context.Context, word string) (Words, error) {
	words := Words{}
	results, err := datamuse.New().Words().RelatedSynonyms(word).Get()
	if err != nil {
		return words, err
	}
	words.fromResults(results)
	return words, nil
}

func (s *Search) Suggestion(ctx context.Context, keyword string) (Words, error) {
	words := Words{}
	results, err := datamuse.New().Suggestions(keyword).Get()
	if err != nil {
		return words, err
	}
	words.fromResults(results)
	return words, nil
}
