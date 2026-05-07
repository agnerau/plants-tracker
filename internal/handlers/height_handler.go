package handlers

import (
	"goplants/internal"
	"goplants/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HeightHandler struct {
	Service *services.HeightService
}

func (h *HeightHandler) CreateHeight(c *gin.Context) {
	var height internal.Height

	if err := c.ShouldBindJSON(&height); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plantID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid plant id"})
		return
	}

	height.PlantID = plantID

	err = h.Service.CreateHeight(&height)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *HeightHandler) GetHeights(c *gin.Context) {
	plantID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid plant id"})
		return
	}
	heights, err := h.Service.GetHeights(plantID, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, heights)
}
