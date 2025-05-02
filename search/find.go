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

	// ray tfidf output to to get most significant word
	wordPriority, err := pipelineWordPriority.Transform(query)
	if err != nil {
		fmt.Printf("Failed to process documents because %v", err)
		return
	}

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

	// reverse the vocab to id : word
	vocab := vectoriser.Vocabulary
	reverseVocab := make(map[int]string)
	for word, id := range vocab {
		reverseVocab[id] = word
	}

	// get top word
	topWord := reverseVocab[maxRow]

	//docs for top word
	docsToTest := keywordIdMap[topWord]

	// calculate and store similarity
	type scoredSentence struct {
		sentence string
		score    float64
	}
	var scoredDocs []scoredSentence
	for _, val := range docsToTest {
		similarity := pairwise.CosineSimilarity(queryVector.(mat.ColViewer).ColView(0), matrix.(mat.ColViewer).ColView(val))
		fmt.Println(corpus[val])
		scoredDocs = append(scoredDocs, scoredSentence{corpus[val], similarity})
	}

	// Sort by highest score
	sort.Slice(scoredDocs, func(i, j int) bool {
		return scoredDocs[i].score > scoredDocs[j].score
	})

	// take top 2 results
	for _, doc := range scoredDocs {

		if len(*result) == 2 {
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
