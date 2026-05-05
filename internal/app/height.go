package app

import (
	"database/sql"
	"goplants/internal/handlers"
	"goplants/internal/repositories"
	"goplants/internal/services"
)

func InitHeightHandler(db *sql.DB) *handlers.HeightHandler {
	repo := &repositories.HeightRepository{DB: db}
	service := &services.HeightService{Repo: repo}
	return &handlers.HeightHandler{Service: service}
}
