package repositories

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetPlants_Repo(t *testing.T) {
	db := setupTestDB(t)
	repo := &PlantRepository{DB: db}

	_, _ = db.Exec("INSERT INTO plants (name) VALUES ('Aloe')")

	plants, err := repo.GetPlants()

	if err != nil {
		t.Fatal(err)
	}

	if len(plants) != 1 {
		t.Errorf("expected 1 plant, got %d", len(plants))
	}
}

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}

	createTable := `
	CREATE TABLE plants (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		nickname TEXT,
		bought_at DATETIME,
		planted_at DATETIME,
		died_at DATETIME,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err = db.Exec(createTable)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	return db
}
