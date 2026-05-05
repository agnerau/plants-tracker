package app

import (
	"database/sql"
	"goplants/internal/handlers"
	"goplants/internal/repositories"
	"goplants/internal/services"
)

func InitPlantHandler(db *sql.DB) *handlers.PlantHandler {
	repo := &repositories.PlantRepository{DB: db}
	service := &services.PlantService{Repo: repo}
	return &handlers.PlantHandler{Service: service}
}
