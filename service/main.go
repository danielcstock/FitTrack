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

	c.IndentedJSON(http.StatusFound, series)
}

func postExercise(c *gin.Context) {
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

func updateExercise(c *gin.Context) {
	var newExercise exercise

	if err := c.BindJSON(&newExercise); err != nil {
		return
	}

	db, err := gorm.Open(sqlite.Open("fittrack.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Model(&newExercise).Updates(exercise{
		Name:   newExercise.Name,
		Times:  newExercise.Times,
		Weight: newExercise.Weight,
	})

	c.IndentedJSON(http.StatusOK, newExercise)
}

func deleteExercise(c *gin.Context) {
	db, err := gorm.Open(sqlite.Open("fittrack.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	id := c.Query("id")
	var exercise exercise
	db.First(&exercise, "id = ?", id)

	db.Delete(&exercise)

	c.IndentedJSON(http.StatusOK, nil)
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
	router.POST("/exercise", postExercise)
	router.PUT("/exercise", updateExercise)
	router.DELETE("/exercise", deleteExercise)

	router.Run("localhost:8081")
}
