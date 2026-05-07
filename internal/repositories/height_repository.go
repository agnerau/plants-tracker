package repositories

import (
	"database/sql"
	"goplants/internal"
	"time"
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
		var createdAtStr string

		err = rows.Scan(&h.ID, &h.Value, &createdAtStr)
		if err != nil {
			return nil, err
		}

		h.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return nil, err
		}
		heights = append(heights, h)
	}

	return heights, nil
}
