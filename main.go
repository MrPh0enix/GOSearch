package main

import (
	"net/http"

	"github.com/MrPh0enix/GOSearch/search"
	"github.com/gin-gonic/gin"
)

type query struct {
	Query string `json:"query"`
}

func findSimilar(c *gin.Context) {
	var inputQuery query

	if err := c.BindJSON(&inputQuery); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := search.MostSimilar(inputQuery.Query)

	c.IndentedJSON(http.StatusOK, result)

}

func main() {
	router := gin.Default()
	router.POST("/calc-similar", findSimilar)
	router.Run("localhost:8080")
}
