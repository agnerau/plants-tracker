package internal

type PlantService struct {
	Repo *PlantRepository
}

func (s *PlantService) CreatePlant(p *Plant) error {
	return s.Repo.CreatePlant(p)
}

func (s *PlantService) GetPlants() ([]Plant, error) {
	return s.Repo.GetPlants()
}
