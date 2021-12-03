package autocomplete

import "github.com/ryeoman/go_dependency_injection/internal/repository/search"

type Suggestion struct {
	Word string `json:"word"`
}

type Suggestions []Suggestion

func (s *Suggestions) fromWords(words search.Words) {
	for i := range words {
		word := words[i]
		*s = append(*s, Suggestion{
			Word: word.Word,
		})
	}
}
