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

	wordPriority, err := pipelineWordPriority.Transform(query)
	if err != nil {
		fmt.Printf("Failed to process documents because %v", err)
		return
	}
	fmt.Printf("Matrix:\n%v\n", mat.Formatted(wordPriority))

	vocab := vectoriser.Vocabulary
	reverseVocab := make(map[int]string)
	for word, id := range vocab {
		reverseVocab[id] = word
	}
	fmt.Println(vocab)

	// Find the max value
	rows, _ := wordPriority.Dims()
	maxVal := wordPriority.At(0, 0)
	maxRow := 0
	for i := 1; i < rows; i++ {
		val := wordPriority.At(i, 0)
		if val > maxVal {
			maxVal = val
			maxRow = i
		}
	}
	fmt.Printf("Max value: %.2f at row %d\n", maxVal, maxRow)

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

		if len(*result) == 1 {
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
