package services

import (
	"goplants/internal"
	"goplants/internal/repositories"
)

type PlantService struct {
	Repo *repositories.PlantRepository
}

func (s *PlantService) CreatePlant(p *internal.Plant) error {
	return s.Repo.CreatePlant(p)
}

func (s *PlantService) GetPlants() ([]internal.Plant, error) {
	return s.Repo.GetPlants()
}
