package service

import (
	"energy_storage/internal/repository"
	"time"
)

type WindTurbineService struct {
	repo repository.WindTurbine
}

func NewWindTurbine(repo repository.WindTurbine) *WindTurbineService {
	return &WindTurbineService{repo: repo}
}

type WTAllData struct {
	Power      float64   `db:"power"`
	Voltage    float64   `db:"voltage"`
	RecordedAt time.Time `db:"recorded_at"`
	Model      string    `db:"model_name"`
	Brand      string    `db:"brand_name"`
}

func (s *WindTurbineService) AllData() ([]*WTAllData, error) {
	data, err := s.repo.AllData()
	if err != nil {
		return nil, err
	}

	newData := make([]*WTAllData, len(data))
	for i, d := range data {
		newData[i] = &WTAllData{
			Power:      d.Power.Float64,
			Voltage:    d.Voltage.Float64,
			RecordedAt: d.RecordedAt,
			Model:      d.Model.String,
			Brand:      d.Brand.String,
		}
	}

	return newData, nil
}
