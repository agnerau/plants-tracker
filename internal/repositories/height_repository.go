package repositories

import (
	"database/sql"
	"goplants/internal"
)

type HeightRepository struct {
	DB *sql.DB
}

func (r *HeightRepository) CreateHeight(h *internal.Height) error {
	query := "INSERT INTO heights(value, plant_id) VALUES (?,?)"
	_, err := r.DB.Exec(query, h.Value, h.PlantID)
	return err
}

func (r *HeightRepository) GetHeights(plantID int) ([]internal.Height, error) {
	query := "SELECT id, value, created_at FROM heights WHERE plant_id = (?)"
	rows, err := r.DB.Query(query, plantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var heights []internal.Height

	for rows.Next() {
		h := internal.Height{PlantID: plantID}
		rows.Scan(&h.ID, &h.Value, &h.CreatedAt)
		heights = append(heights, h)
	}

	return heights, nil
}
