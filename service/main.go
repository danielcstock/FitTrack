package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type exercise struct {
	Name   string    `json: "name"`
	Times  int       `json: "times"`
	Weight int       `json: "weight"`
	Date   time.Time `json: "date"`
}

var series = []exercise{
	{Name: "Agachamento livre", Times: 12, Weight: 14, Date: time.Now().UTC()},
	{Name: "Agachamento", Times: 12, Weight: 14, Date: time.Now().UTC()},
}

func getExercises(c *gin.Context) {
	file, _ := os.Open("track.json")
	b, _ := io.ReadAll(file)

	var sb strings.Builder
	sb.WriteString("[")
	sb.WriteString(string(b))
	sb.WriteString("]")

	json.Unmarshal([]byte(sb.String()), &series)

	c.IndentedJSON(http.StatusOK, series)
}

func postExercises(c *gin.Context) {
	var newExercise exercise

	if err := c.BindJSON(&newExercise); err != nil {
		return
	}

	newExercise.Date = time.Now().UTC()

	b, err := json.Marshal(newExercise)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	f, err := os.OpenFile("track.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var sb strings.Builder
	sb.WriteString(",\n")
	sb.WriteString(string(b))

	if _, err = f.WriteString(sb.String()); err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusCreated, newExercise)
}

func main() {
	router := gin.Default()

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"PUT", "PATCH"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))
	router.Use(cors.Default())

	router.GET("/exercise", getExercises)
	router.POST("/exercise", postExercises)

	router.Run("localhost:8080")
}
