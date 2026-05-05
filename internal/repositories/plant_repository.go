package repositories

import (
	"database/sql"
	"goplants/internal"
)

type PlantRepository struct {
	DB *sql.DB
}

func (r *PlantRepository) CreatePlant(p *internal.Plant) error {
	query := "INSERT INTO plants(name, nickname) VALUES (?, ?)"
	_, err := r.DB.Exec(query, p.Name, p.Nickname)
	return err
}

func (r *PlantRepository) GetPlants() ([]internal.Plant, error) {
	rows, err := r.DB.Query("SELECT id, name, nickname FROM plants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plants []internal.Plant

	for rows.Next() {
		var p internal.Plant
		rows.Scan(&p.ID, &p.Name, &p.Nickname)
		plants = append(plants, p)
	}

	return plants, nil
}
