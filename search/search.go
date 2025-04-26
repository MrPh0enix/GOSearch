package search

import (
	"fmt"

	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"gonum.org/v1/gonum/mat"
)

func TfIdfSearch(corpus []string, query string, result *string) {
	vectoriser := nlp.NewCountVectoriser()
	tfIdf := nlp.NewTfidfTransformer()
	reducer := nlp.NewTruncatedSVD(100) //Reduce to 100 features
	pipeline := nlp.NewPipeline(vectoriser, tfIdf, reducer)

	matrix, err := pipeline.FitTransform(corpus...)
	if err != nil {
		fmt.Printf("Failed to process documents because %v", err)
		return
	}

	queryVector, err := pipeline.Transform(query)
	if err != nil {
		fmt.Printf("Failed to process documents because %v", err)
		return
	}

	highestSimilarity := -1.0
	var highIdx int
	_, docs := matrix.Dims() // Columns represent the documents in the corpus
	for i := 0; i < docs; i++ {
		similarity := pairwise.CosineSimilarity(queryVector.(mat.ColViewer).ColView(0), matrix.(mat.ColViewer).ColView(i))
		if similarity > highestSimilarity {
			highIdx = i
			highestSimilarity = similarity
		}
	}

	*result = corpus[highIdx]
}

func MostSimilar(query string) string {
	docs := []string{
		"the quick brown fox jumped over the lazy dog",
		"the quick brown fox",
		"the dog jumped over the quick fox",
	}

	// query := "the brown fox ran around the dog"
	var result string
	TfIdfSearch(docs, query, &result)
	return result
}
