package services

import (
	"goplants/internal"
)

type PlantRepository interface {
	GetPlants() ([]internal.Plant, error)
	GetPlant(plantID int) (internal.Plant, error)
	CreatePlant(p *internal.Plant) error
}
type PlantService struct {
	Repo PlantRepository
}

func (s *PlantService) CreatePlant(p *internal.Plant) error {
	return s.Repo.CreatePlant(p)
}

func (s *PlantService) GetPlants() ([]internal.Plant, error) {
	return s.Repo.GetPlants()
}
func (s *PlantService) GetPlant(plantID int) (internal.Plant, error) {
	return s.Repo.GetPlant(plantID)
}
