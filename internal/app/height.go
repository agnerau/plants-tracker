package app

import (
	"database/sql"
	"goplants/internal/handlers"
	"goplants/internal/repositories"
	"goplants/internal/services"

	"github.com/redis/go-redis/v9"
)

func InitHeightHandler(db *sql.DB, rdb *redis.Client) *handlers.HeightHandler {
	repo := &repositories.HeightRepository{DB: db}
	service := &services.HeightService{Repo: repo, RDB: rdb}
	return &handlers.HeightHandler{Service: service}
}
