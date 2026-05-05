package services

import (
	"goplants/internal"
	"goplants/internal/repositories"
)

type HeightService struct {
	Repo *repositories.HeightRepository
}

func (s *HeightService) CreateHeight(h *internal.Height) error {
	return s.Repo.CreateHeight(h)
}

func (s *HeightService) GetHeights(plantID int) ([]internal.Height, error) {
	return s.Repo.GetHeights(plantID)
}
