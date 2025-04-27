package search

import (
	"fmt"

	"github.com/james-bowman/nlp"
	"gonum.org/v1/gonum/mat"
)

var keywordIdMap = make(map[string][]int)

// docs are populated with default values which can be removed in actual use.
var corpus = []string{
	"The cat sat on the mat",
	"A dog chased the cat",
	"Birds fly high in the blue sky",
	"The quick brown fox jumps over the lazy dog",
}
var pipeline *nlp.Pipeline
var matrix mat.Matrix

func init() {
	vectoriser := nlp.NewCountVectoriser()
	tfIdf := nlp.NewTfidfTransformer()
	reducer := nlp.NewTruncatedSVD(100) //Reduce to 100 features
	pipeline = nlp.NewPipeline(vectoriser, tfIdf, reducer)

	// Fit the data
	var err error
	matrix, err = pipeline.FitTransform(corpus...)
	if err != nil {
		fmt.Printf("Failed to process documents because %v", err)
		return
	}
}

func AddDoc(inputDoc string) {
	corpus = append(corpus, inputDoc)

	// Fit the data
	var err error
	matrix, err = pipeline.FitTransform(corpus...)
	if err != nil {
		fmt.Printf("Failed to process documents because %v", err)
		return
	}

}
