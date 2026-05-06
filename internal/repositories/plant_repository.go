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
		if err = rows.Scan(&p.ID, &p.Name, &p.Nickname); err != nil {
			return nil, err
		}
		plants = append(plants, p)
	}

	return plants, nil
}

func (r *PlantRepository) GetPlant(plantID int) (internal.Plant, error) {
	plant := internal.Plant{}
	query := ("SELECT name, nickname, bought_at, planted_at, died_at FROM plants WHERE id = ?")
	row := r.DB.QueryRow(query, plantID)
	err := row.Scan(
		&plant.Name,
		&plant.Nickname,
		&plant.BoughtAt,
		&plant.PlantedAt,
		&plant.DiedAt,
	)

	if err != nil {
		return plant, err
	}

	return plant, nil
}
