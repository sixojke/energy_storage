package service

import (
	"energy_storage/internal/repository"
	"time"
)

type SolarPanelService struct {
	repo repository.SolarPanel
}

func NewSolarPanel(repo repository.SolarPanel) *SolarPanelService {
	return &SolarPanelService{repo: repo}
}

type SPAllData struct {
	NominalPower float64   `json:"nominal_power"`
	RecordedAt   time.Time `json:"recorded_at"`
	Model        string    `json:"model_name"`
	Brand        string    `json:"brand_name"`
	Length       float64   `json:"length"`
	Width        float64   `json:"width"`
	Weight       float64   `json:"weight"`
	PanelCount   int       `json:"panel_count"`
	BatteryCount int       `json:"battery_count"`
	Voltage      float64   `json:"voltage"`
}

func (s *SolarPanelService) AllData() ([]*SPAllData, error) {
	data, err := s.repo.AllData()
	if err != nil {
		return nil, err
	}

	newData := make([]*SPAllData, len(data))
	for i, d := range data {
		newData[i] = &SPAllData{
			NominalPower: d.NominalPower.Float64,
			RecordedAt:   d.RecordedAt,
			Model:        d.Model.String,
			Brand:        d.Brand.String,
			Length:       d.Length.Float64,
			Width:        d.Width.Float64,
			Weight:       d.Weight.Float64,
			PanelCount:   int(d.PanelCount.Int64),
			BatteryCount: int(d.BatteryCount.Int64),
			Voltage:      d.Voltage.Float64,
		}
	}

	return newData, nil
}
