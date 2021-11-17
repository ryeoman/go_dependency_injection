package similarity

import "github.com/ryeoman/go_dependency_injection/internal/repository/search"

type Similar struct {
	Word string `json:"word"`
}

type Similars []Similar

func (s *Similars) fromWords(words search.Words) {
	for i, _ := range words {
		word := words[i]
		*s = append(*s, Similar{
			Word: word.Word,
		})
	}
}
