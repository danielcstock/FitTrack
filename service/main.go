package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type exercise struct {
	Name   string `json: "name"`
	Times  int    `json: "times"`
	Weight int    `json: "weight"`
}

var series = []exercise{
	{Name: "Agachamento livre", Times: 12, Weight: 14},
	{Name: "Agachamento", Times: 12, Weight: 14},
}

func getExercises(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, series)
}

func postExercises(c *gin.Context) {
	var newExercise exercise

	if err := c.BindJSON(&newExercise); err != nil {
		return
	}

	series = append(series, newExercise)
	c.IndentedJSON(http.StatusCreated, newExercise)
}

func main() {
	router := gin.Default()
	router.GET("/exercise", getExercises)
	router.POST("/exercise", postExercises)

	router.Run("localhost:8080")
}
