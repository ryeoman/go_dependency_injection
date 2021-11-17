package search

import "github.com/kostaspt/go-datamuse/v2"

type Word struct {
	Word  string
	Score int
}

type Words []Word

func (w *Words) fromResults(results datamuse.Results) {
	for i, _ := range results {
		result := results[i]
		word := Word{
			Word:  result.Word,
			Score: result.Score,
		}
		*w = append(*w, word)
	}
}
