package services

import (
	"context"
	"encoding/json"
	"goplants/internal"
	"time"

	"github.com/redis/go-redis/v9"
)

type PlantRepository interface {
	GetPlants() ([]internal.Plant, error)
	GetPlant(plantID int) (internal.Plant, error)
	CreatePlant(p *internal.Plant) error
}
type PlantService struct {
	Repo PlantRepository
	RDB  *redis.Client
}

func (s *PlantService) CreatePlant(p *internal.Plant) error {
	err := s.Repo.CreatePlant(p)
	if err != nil {
		return err
	}

	s.RDB.Del(context.Background(), "plants:all")

	return nil
}

func (s *PlantService) GetPlants(ctx context.Context) ([]internal.Plant, error) {
	cached, err := s.RDB.Get(ctx, "plants:all").Result()
	if err == nil {
		var plants []internal.Plant
		json.Unmarshal([]byte(cached), &plants)
		return plants, nil
	}

	plants, err := s.Repo.GetPlants()
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(plants)
	s.RDB.Set(ctx, "plants:all", data, time.Hour)

	return plants, nil
}
func (s *PlantService) GetPlant(plantID int) (internal.Plant, error) {
	return s.Repo.GetPlant(plantID)
}
