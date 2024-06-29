package service

import "energy_storage/internal/repository"

type ElectrochemicalBattery interface {
	AllData() ([]*EBAllDataOut, error)
}

type ThermalStorage interface {
	AllData() ([]*TSAllData, error)
}

type HydrogenStorage interface {
	AllData() ([]*HSAllData, error)
}

type WindTurbine interface {
	AllData() ([]*WTAllData, error)
}

type SolarPanel interface {
	AllData() ([]*SPAllData, error)
}

type Deps struct {
	Repo repository.Repository
}

type Service struct {
	ElectrochemicalBattery
	ThermalStorage
	HydrogenStorage
	WindTurbine
	SolarPanel
}

func NewSerivice(deps *Deps) *Service {
	return &Service{
		ElectrochemicalBattery: NewElectrochemicalBatterySerive(deps.Repo.ElectrochemicalBattery),
		ThermalStorage:         NewThermalStorage(deps.Repo.ThermalStorage),
		HydrogenStorage:        NewHydrogenStorage(deps.Repo.HydrogenStorage),
		WindTurbine:            NewWindTurbine(deps.Repo.WindTurbine),
		SolarPanel:             NewSolarPanel(deps.Repo.SolarPanel),
	}
}
