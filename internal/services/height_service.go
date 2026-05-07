package services

import (
	"context"
	"encoding/json"
	"goplants/internal"
	"goplants/internal/repositories"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type HeightService struct {
	Repo *repositories.HeightRepository
	RDB  *redis.Client
}

func (s *HeightService) CreateHeight(h *internal.Height) error {
	err := s.Repo.CreateHeight(h)
	if err != nil {
		return err
	}

	s.RDB.Del(context.Background(), "heights:"+strconv.Itoa(h.PlantID))
	return nil
}

func (s *HeightService) GetHeights(plantID int, ctx context.Context) ([]internal.Height, error) {
	cached, err := s.RDB.Get(ctx, "heights:"+strconv.Itoa(plantID)).Result()
	if err == nil {
		var heights []internal.Height
		json.Unmarshal([]byte(cached), &heights)
		return heights, nil
	}

	heights, err := s.Repo.GetHeights(plantID)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(heights)
	s.RDB.Set(ctx, "heights:"+strconv.Itoa(plantID), data, time.Hour)

	return heights, nil
}
