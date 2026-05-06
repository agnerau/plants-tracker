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

	r.POST("api/plants", plantHandler.CreatePlant)
	r.GET("api/plants", plantHandler.GetPlants)
	r.GET("api/plants/:id", plantHandler.GetPlant)
	r.POST("api/plants/:id/heights", heightHandler.CreateHeight)
	r.GET("api/plants/:id/heights", heightHandler.GetHeights)

	r.Static("/static", "./web")
	r.GET("/", func(c *gin.Context) {
		c.File("./web/index.html")
	})
	r.GET("/plants/:id", func(c *gin.Context) {
		c.File("./web/plant.html")
	})

	r.Run(":8080")
}
