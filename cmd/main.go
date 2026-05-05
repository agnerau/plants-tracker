package main

import (
	"goplants/internal"
	"goplants/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()

	repo := &internal.PlantRepository{DB: database}
	service := &internal.PlantService{Repo: repo}
	handler := &internal.PlantHandler{Service: service}

	r := gin.Default()

	r.POST("/plants", handler.CreatePlant)
	r.GET("/plants", handler.GetPlants)

	r.Run(":8080")
}
