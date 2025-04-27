package main

import (
	"net/http"

	"github.com/MrPh0enix/GOSearch/search"
	"github.com/gin-gonic/gin"
)

type query struct {
	Query string `json:"query"`
}

func findSimilarHandler(c *gin.Context) {
	var inputQuery query

	if err := c.BindJSON(&inputQuery); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := search.MostSimilar(inputQuery.Query)

	c.IndentedJSON(http.StatusOK, result)

}

func addDocHandler(c *gin.Context) {
	var inputQuery query

	if err := c.BindJSON(&inputQuery); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	search.AddDoc(inputQuery.Query)

	c.IndentedJSON(http.StatusOK, "Added document")
}

func main() {
	router := gin.Default()
	router.POST("/calc-similar", findSimilarHandler) //returns top 10 docs
	router.POST("/add-doc", addDocHandler)           // adds a new doc by calculating a keyword
	router.Run("localhost:8080")
}
