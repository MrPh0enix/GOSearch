package search

import (
	"bufio"
	"fmt"
	"os"
)

func loadExamples() *[]string {
	file, err := os.Open("sentences.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer file.Close()

	var exampleDocs []string

	// read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		exampleDocs = append(exampleDocs, scanner.Text())
	}

	return &exampleDocs
}
