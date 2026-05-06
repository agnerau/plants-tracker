package handlers

import (
	"goplants/internal"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlantService interface {
	GetPlants() ([]internal.Plant, error)
	GetPlant(plantID int) (internal.Plant, error)
	CreatePlant(p *internal.Plant) error
}

type PlantHandler struct {
	Service PlantService
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

func (h *PlantHandler) GetPlant(c *gin.Context) {
	plantID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid plant id"})
		return
	}
	plant, err := h.Service.GetPlant(plantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, plant)
}
