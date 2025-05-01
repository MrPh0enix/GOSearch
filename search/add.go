package search

import (
	"fmt"
	"strings"

	"github.com/bbalet/stopwords"
	"github.com/james-bowman/nlp"
	"gonum.org/v1/gonum/mat"
)

var keywordIdMap = make(map[string][]int)

// docs are populated with default values which can be removed in actual use.
var corpus []string
var pipeline *nlp.Pipeline
var pipelineWordPriority *nlp.Pipeline
var vectoriser *nlp.CountVectoriser
var matrix mat.Matrix

func init() {
	vectoriser = nlp.NewCountVectoriser()
	tfIdf := nlp.NewTfidfTransformer()
	reducer := nlp.NewTruncatedSVD(100) //Reduce to 100 features
	pipeline = nlp.NewPipeline(vectoriser, tfIdf, reducer)
	pipelineWordPriority = nlp.NewPipeline(vectoriser, tfIdf)

	// example docs, delete after testing
	var exampleDocs = []string{
		"The cat sat on the mat",
		"A dog chased the cat",
		"Birds fly high in the blue sky",
		"The quick brown fox jumps over the lazy dog",
	}

	for _, doc := range exampleDocs {
		AddDoc(doc)
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

	cleanDoc := stopwords.CleanString(inputDoc, "en", false)
	cleanDocLis := strings.Fields(cleanDoc)

	for _, wrd := range cleanDocLis {
		keywordIdMap[wrd] = append(keywordIdMap[wrd], len(corpus)-1)
	}

}
