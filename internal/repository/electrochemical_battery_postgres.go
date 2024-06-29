package repository

import (
	"database/sql"
	"energy_storage/internal/domain"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type ElectrochemicalBatteryPostgres struct {
	db *sqlx.DB
}

func NewElectrochemicalBatteryPostgres(db *sqlx.DB) *ElectrochemicalBatteryPostgres {
	return &ElectrochemicalBatteryPostgres{db: db}
}

type EBAllDataOut struct {
	TemperatureRange     domain.Int4Range `db:"temperature_range"`
	InputVoltage         sql.NullFloat64  `db:"input_voltage"`
	OutputVoltage        sql.NullFloat64  `db:"output_voltage"`
	InternalResistance   sql.NullFloat64  `db:"internal_resistance"`
	OperatingTemperature sql.NullFloat64  `db:"operating_temperature"`
	RecordedAt           time.Time        `db:"recorded_at"`
	SpecificEnergy       sql.NullFloat64  `db:"specific_energy"`
	CycleLife            sql.NullInt64    `db:"cycle_life"`
	ChargeTime           domain.Int4Range `db:"charge_time"`
	DischargeTime        domain.Int4Range `db:"discharge_time"`
	Efficiency           sql.NullFloat64  `db:"efficiency"`
	SelfDischarge        domain.NumRange  `db:"self_discharge"`
	Model                string           `db:"model_name"`
}

func (r *ElectrochemicalBatteryPostgres) AllData() ([]*EBAllDataOut, error) {
	query := `
      	SELECT
      		t1.temperature_range,
      		t1.input_voltage,
      		t1.output_voltage,
      		t1.internal_resistance,
      		t1.operating_temperature,
	  		t1.recorded_at,
      		t2.specific_energy,
      		t2.cycle_life,
      		t2.charge_time,
      		t2.discharge_time, 
      		t2.efficiency,
      		t2.self_discharge,
      		t3.model_name
    	FROM electrochemical_battery_history t1
    	INNER JOIN energy_storage_characteristics t2 ON t1.characteristics_id = t2.id
    	INNER JOIN model t3 ON t1.model_id = t3.id
    `

	var characteristics []*EBAllDataOut
	if err := r.db.Select(&characteristics, query); err != nil {
		return nil, fmt.Errorf("failed to get electrochemical battery data: %v", err)
	}

	return characteristics, nil
}
