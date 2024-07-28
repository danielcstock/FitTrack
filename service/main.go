package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type exercise struct {
	gorm.Model
	Name   string `json: "name"`
	Times  int    `json: "times"`
	Weight int    `json: "weight"`
}

func getExercises(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("fittrack.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var series []exercise
	db.Find(&series).Order("created_at desc")

	c.IndentedJSON(http.StatusOK, series)
}

func getExerciseByName(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("fittrack.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	name := c.Query("name")
	var series exercise
	db.First(&series, "name like ?", "%"+name+"%").Order("created_at desc")

	c.IndentedJSON(http.StatusOK, series)
}

func postExercises(c *gin.Context) {
	var newExercise exercise

	if err := c.BindJSON(&newExercise); err != nil {
		return
	}

	db, err := gorm.Open(sqlite.Open("fittrack.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&exercise{
		Name:   newExercise.Name,
		Times:  newExercise.Times,
		Weight: newExercise.Weight,
	})

	c.IndentedJSON(http.StatusCreated, newExercise)
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	db, err := gorm.Open(sqlite.Open("fittrack.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&exercise{})

	router.GET("/exercises", getExercises)
	router.GET("/exercise", getExerciseByName)
	router.POST("/exercise", postExercises)

	router.Run("localhost:8081")
}
