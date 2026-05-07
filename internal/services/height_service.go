package services

import (
	"goplants/internal"
	"goplants/internal/repositories"

	"github.com/redis/go-redis/v9"
)

type HeightService struct {
	Repo *repositories.HeightRepository
	RDB  *redis.Client
}

func (s *HeightService) CreateHeight(h *internal.Height) error {
	return s.Repo.CreateHeight(h)
}

func (s *HeightService) GetHeights(plantID int) ([]internal.Height, error) {
	return s.Repo.GetHeights(plantID)
}
