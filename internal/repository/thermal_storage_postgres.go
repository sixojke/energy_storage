package repository

import (
	"database/sql"
	"energy_storage/internal/domain"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type ThermalStoragePostgres struct {
	db *sqlx.DB
}

func NewThermalStoragePostgres(db *sqlx.DB) *ThermalStoragePostgres {
	return &ThermalStoragePostgres{db: db}
}

type TSAllData struct {
	ThermalPower     sql.NullFloat64  `db:"thermal_power"`
	Efficiency       sql.NullFloat64  `db:"efficiency"`
	RecordedAt       time.Time        `db:"recorded_at"`
	SpecificEnergy   sql.NullFloat64  `db:"specific_energy"`
	CycleLife        sql.NullInt64    `db:"cycle_life"`
	TemperatureRange domain.Int4Range `db:"temperature_range"`
	SelfDischarge    domain.NumRange  `db:"self_discharge"`
}

func (r *ThermalStoragePostgres) AllData() ([]*TSAllData, error) {
	query := `
		SELECT 
			t1.thermal_power,
			t1.efficiency,
			t1.recorded_at,
			t2.specific_energy,
			t2.cycle_life,
			t2.temperature_range,
			t2.self_discharge
		FROM thermal_storage_history t1
		INNER JOIN energy_storage_characteristics t2 ON t1.characteristics_id = t2.id
	`

	var characteristics []*TSAllData
	if err := r.db.Select(&characteristics, query); err != nil {
		return nil, fmt.Errorf("failed to get thermal storage data: %v", err)
	}

	return characteristics, nil
}
