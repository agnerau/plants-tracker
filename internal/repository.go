package internal

import (
	"database/sql"
)

type PlantRepository struct {
	DB *sql.DB
}

func (r *PlantRepository) CreatePlant(p *Plant) error {
	query := "INSERT INTO plants(name, nickname) VALUES (?, ?)"
	_, err := r.DB.Exec(query, p.Name, p.Nickname)
	return err
}

func (r *PlantRepository) GetPlants() ([]Plant, error) {
	rows, err := r.DB.Query("SELECT id, name, nickname FROM plants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var plants []Plant

	for rows.Next() {
		var p Plant
		rows.Scan(&p.ID, &p.Name, &p.Nickname)
		plants = append(plants, p)
	}

	return plants, nil
}
