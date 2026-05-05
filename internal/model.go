package internal

import "time"

type Plant struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Nickname  *string    `json:"nickname"`
	BoughtAt  *time.Time `json:"bought_at"`
	PlantedAt *time.Time `json:"planted_at"`
	DiedAt    *time.Time `json:"died_at"`
	CreatedAt time.Time  `json:"created_at"`
}

type Height struct {
	ID        int       `json:"id"`
	Value     float64   `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	PlantID   int       `json:"plant_id"`
}
