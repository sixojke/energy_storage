package repository

import "github.com/jmoiron/sqlx"

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
	Postgres *sqlx.DB
}

type Repository struct {
	ElectrochemicalBattery
	ThermalStorage
	HydrogenStorage
	WindTurbine
	SolarPanel
}

func NewRepository(deps *Deps) *Repository {
	return &Repository{
		ElectrochemicalBattery: NewElectrochemicalBatteryPostgres(deps.Postgres),
		ThermalStorage:         NewThermalStoragePostgres(deps.Postgres),
		HydrogenStorage:        NewHydrogenStoragePostgres(deps.Postgres),
		WindTurbine:            NewWindTurbinePostgres(deps.Postgres),
		SolarPanel:             NewSolarPanelPostgres(deps.Postgres),
	}
}
