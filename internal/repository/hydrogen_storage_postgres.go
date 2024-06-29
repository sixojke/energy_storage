package repository

import (
	"database/sql"
	"energy_storage/internal/domain"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type HydrogenStoragePostgres struct {
	db *sqlx.DB
}

func NewHydrogenStoragePostgres(db *sqlx.DB) *HydrogenStoragePostgres {
	return &HydrogenStoragePostgres{db: db}
}

type HSAllData struct {
	Id               int              `db:"id"`
	TemperatureRange domain.Int4Range `db:"temperature_range"`
	NominalVoltage   sql.NullFloat64  `db:"nominal_voltage"`
	RecordedAt       time.Time        `db:"recorded_at"`
	BrandName        sql.NullString   `db:"brand_name"`
	ModelName        sql.NullString   `db:"model_name"`
	SpecificEnergy   sql.NullFloat64  `db:"specific_energy"`
	Efficiency       sql.NullFloat64  `db:"efficiency"`
}

func (r *HydrogenStoragePostgres) AllData() ([]*HSAllData, error) {
	query := `
		SELECT
			t1.id,
			t1.temperature_range,
			t1.nominal_voltage,
			t1.recorded_at,
			t2.brand_name,
			t3.model_name,
			t4.specific_energy,
			t4.efficiency
			FROM hydrogen_storage_history t1 
		INNER JOIN brand t2 ON t1.brand_id = t2.id
		INNER JOIN model t3 ON t1.model_id = t3.id
		INNER JOIN energy_storage_characteristics t4 ON t1.characteristics_id = t4.id
	`
	var characteristics []*HSAllData
	if err := r.db.Select(&characteristics, query); err != nil {
		return nil, fmt.Errorf("failed to get hydrogen storage data: %v", err)
	}

	return characteristics, nil
}
