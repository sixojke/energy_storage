package service

import (
	"energy_storage/internal/domain"
	"energy_storage/internal/repository"
	"time"
)

type ElectrochemicalBatterySerive struct {
	repo repository.ElectrochemicalBattery
}

func NewElectrochemicalBatterySerive(repo repository.ElectrochemicalBattery) *ElectrochemicalBatterySerive {
	return &ElectrochemicalBatterySerive{
		repo: repo,
	}
}

type EBAllDataOut struct {
	TemperatureRange     domain.Int4Range `json:"temperature_range"`
	InputVoltage         float64          `json:"input_voltage"`
	OutputVoltage        float64          `json:"output_voltage"`
	InternalResistance   float64          `json:"internal_resistance"`
	OperatingTemperature float64          `json:"operating_temperature"`
	RecordedAt           time.Time        `json:"recorded_at"`
	SpecificEnergy       float64          `json:"specific_energy"`
	CycleLife            int              `json:"cycle_life"`
	ChargeTime           domain.Int4Range `json:"charge_time"`
	DischargeTime        domain.Int4Range `json:"discharge_time"`
	Efficiency           float64          `json:"efficiency"`
	SelfDischarge        domain.NumRange  `json:"self_discharge"`
	Model                string           `json:"model_name"`
}

func (s *ElectrochemicalBatterySerive) AllData() ([]*EBAllDataOut, error) {
	data, err := s.repo.AllData()
	if err != nil {
		return nil, err
	}

	newData := make([]*EBAllDataOut, len(data))
	for i, d := range data {
		newData[i] = &EBAllDataOut{
			TemperatureRange:     d.TemperatureRange,
			InputVoltage:         d.InputVoltage.Float64,
			OutputVoltage:        d.OutputVoltage.Float64,
			InternalResistance:   d.InternalResistance.Float64,
			OperatingTemperature: d.OperatingTemperature.Float64,
			RecordedAt:           d.RecordedAt,
			SpecificEnergy:       d.SpecificEnergy.Float64,
			CycleLife:            int(d.CycleLife.Int64),
			ChargeTime:           d.ChargeTime,
			DischargeTime:        d.DischargeTime,
			Efficiency:           d.Efficiency.Float64,
			SelfDischarge:        d.SelfDischarge,
			Model:                d.Model,
		}
	}

	return newData, nil
}
