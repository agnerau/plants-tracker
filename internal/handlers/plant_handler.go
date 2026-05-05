package handlers

import (
	"goplants/internal"
	"goplants/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlantHandler struct {
	Service *services.PlantService
}

func (h *PlantHandler) CreatePlant(c *gin.Context) {
	var plant internal.Plant

	if err := c.ShouldBindJSON(&plant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.CreatePlant(&plant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *PlantHandler) GetPlants(c *gin.Context) {
	plants, err := h.Service.GetPlants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plants)
}
