package main

import (
	"goplants/internal/app"
	"goplants/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()

	plantHandler := app.InitPlantHandler(database)
	heightHandler := app.InitHeightHandler(database)

	r := gin.Default()

	r.POST("/plants", plantHandler.CreatePlant)
	r.GET("/plants", plantHandler.GetPlants)
	r.POST("/plants/:id/heights", heightHandler.CreateHeight)
	r.GET("/plants/:id/heights", heightHandler.GetHeights)

	r.Run(":8080")
}
