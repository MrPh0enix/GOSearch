package search

import (
	"fmt"
	"sort"

	"github.com/james-bowman/nlp/measures/pairwise"
	"gonum.org/v1/gonum/mat"
)

func TfIdfSearch(query string, result *[]string) {

	// Convert query to vector
	queryVector, err := pipeline.Transform(query)
	if err != nil {
		fmt.Printf("Failed to process documents because %v", err)
		return
	}

	// calculate and store similarity
	type scoredSentence struct {
		sentence string
		score    float64
	}
	var scoredDocs []scoredSentence
	_, docs := matrix.Dims() // Columns represent the documents in the corpus
	for i := 0; i < docs; i++ {
		similarity := pairwise.CosineSimilarity(queryVector.(mat.ColViewer).ColView(0), matrix.(mat.ColViewer).ColView(i))
		scoredDocs = append(scoredDocs, scoredSentence{corpus[i], similarity})
	}

	// Sort by highest score
	sort.Slice(scoredDocs, func(i, j int) bool {
		return scoredDocs[i].score > scoredDocs[j].score
	})

	// take top 2 results
	for _, doc := range scoredDocs {
		if len(*result) == 10 {
			break
		}
		*result = append(*result, doc.sentence)
	}

}

func MostSimilar(query string) []string {

	var result []string
	TfIdfSearch(query, &result)
	return result
}
