package service

import (
	"energy_storage/internal/domain"
	"energy_storage/internal/repository"
	"time"
)

type HydrogenStorageService struct {
	repo repository.HydrogenStorage
}

func NewHydrogenStorage(repo repository.HydrogenStorage) *HydrogenStorageService {
	return &HydrogenStorageService{repo: repo}
}

type HSAllData struct {
	Id               int              `json:"id"`
	TemperatureRange domain.Int4Range `json:"temperature_range"`
	NominalVoltage   float64          `json:"nominal_voltage"`
	RecordedAt       time.Time        `json:"recorded_at"`
	BrandName        string           `json:"brand_name"`
	ModelName        string           `json:"model_name"`
	SpecificEnergy   float64          `json:"specific_energy"`
	Efficiency       float64          `json:"efficiency"`
}

func (s *HydrogenStorageService) AllData() ([]*HSAllData, error) {
	data, err := s.repo.AllData()
	if err != nil {
		return nil, err
	}

	newData := make([]*HSAllData, len(data))
	for i, d := range data {
		newData[i] = &HSAllData{
			Id:               d.Id,
			TemperatureRange: d.TemperatureRange,
			NominalVoltage:   d.NominalVoltage.Float64,
			RecordedAt:       d.RecordedAt,
			BrandName:        d.BrandName.String,
			ModelName:        d.ModelName.String,
			SpecificEnergy:   d.SpecificEnergy.Float64,
			Efficiency:       d.Efficiency.Float64,
		}
	}

	return newData, nil
}
