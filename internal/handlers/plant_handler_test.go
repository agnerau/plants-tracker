package handlers

import (
	"bytes"
	"errors"
	"goplants/internal"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter(handler *PlantHandler) *gin.Engine {
	r := gin.New()

	r.POST("/plants", handler.CreatePlant)
	r.GET("/plants", handler.GetPlants)

	return r
}

type MockPlantService struct {
	GetFn    func() ([]internal.Plant, error)
	GetOneFn func(plantID int) (internal.Plant, error)
	CreateFn func(p *internal.Plant) error
}

func (m *MockPlantService) GetPlants() ([]internal.Plant, error) {
	return m.GetFn()
}

func (m *MockPlantService) CreatePlant(p *internal.Plant) error {
	return m.CreateFn(p)
}

func (m *MockPlantService) GetPlant(plantID int) (internal.Plant, error) {
	return m.GetOneFn(plantID)
}

func TestGetPlants(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockPlantService{
		GetFn: func() ([]internal.Plant, error) {
			return []internal.Plant{
				{ID: 1, Name: "Aloe"},
			}, nil
		},
	}

	handler := &PlantHandler{Service: mockService}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodGet, "/plants", nil)
	c.Request = req

	handler.GetPlants(c)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestGetPlants_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockPlantService{
		GetFn: func() ([]internal.Plant, error) {
			return nil, errors.New("db error")
		},
	}

	handler := &PlantHandler{Service: mockService}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req := httptest.NewRequest(http.MethodGet, "/plants", nil)
	c.Request = req

	handler.GetPlants(c)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", w.Code)
	}
}

func TestCreatePlant(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockPlantService{

		CreateFn: func(p *internal.Plant) error {
			return nil
		},
	}

	handler := &PlantHandler{Service: mockService}
	router := setupRouter(handler)

	body := `{"name":"Aloe"}`

	req := httptest.NewRequest(http.MethodPost, "/plants", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", w.Code)
	}
}
