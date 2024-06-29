package service

import (
	"energy_storage/internal/domain"
	"energy_storage/internal/repository"
	"time"
)

type ThermalStorageService struct {
	repo repository.ThermalStorage
}

func NewThermalStorage(repo repository.ThermalStorage) *ThermalStorageService {
	return &ThermalStorageService{repo: repo}
}

type TSAllData struct {
	ThermalPower     float64          `json:"thermal_power"`
	Efficiency       float64          `json:"efficiency"`
	RecordedAt       time.Time        `json:"recorded_at"`
	SpecificEnergy   float64          `json:"specific_energy"`
	CycleLife        int              `json:"cycle_life"`
	TemperatureRange domain.Int4Range `json:"temperature_range"`
	SelfDischarge    domain.NumRange  `json:"self_discharge"`
}

func (s *ThermalStorageService) AllData() ([]*TSAllData, error) {
	data, err := s.repo.AllData()
	if err != nil {
		return nil, err
	}

	newData := make([]*TSAllData, len(data))
	for i, d := range data {
		newData[i] = &TSAllData{
			ThermalPower:     d.ThermalPower.Float64,
			Efficiency:       d.Efficiency.Float64,
			RecordedAt:       d.RecordedAt,
			SpecificEnergy:   d.SpecificEnergy.Float64,
			CycleLife:        int(d.CycleLife.Int64),
			TemperatureRange: d.TemperatureRange,
			SelfDischarge:    d.SelfDischarge,
		}
	}

	return newData, nil
}
