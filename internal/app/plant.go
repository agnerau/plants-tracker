package app

import (
	"database/sql"
	"goplants/internal/handlers"
	"goplants/internal/repositories"
	"goplants/internal/services"

	"github.com/redis/go-redis/v9"
)

func InitPlantHandler(db *sql.DB, rdb *redis.Client) *handlers.PlantHandler {
	repo := &repositories.PlantRepository{DB: db}
	service := &services.PlantService{Repo: repo, RDB: rdb}
	return &handlers.PlantHandler{Service: service}
}
