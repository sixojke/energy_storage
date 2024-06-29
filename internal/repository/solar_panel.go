package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type SolarPanelPostgres struct {
	db *sqlx.DB
}

func NewSolarPanelPostgres(db *sqlx.DB) *SolarPanelPostgres {
	return &SolarPanelPostgres{db: db}
}

type SPAllData struct {
	NominalPower sql.NullFloat64 `db:"nominal_power"`
	RecordedAt   time.Time       `db:"recorded_at"`
	Model        sql.NullString  `db:"model_name"`
	Brand        sql.NullString  `db:"brand_name"`
	Length       sql.NullFloat64 `db:"length"`
	Width        sql.NullFloat64 `db:"width"`
	Weight       sql.NullFloat64 `db:"weight"`
	PanelCount   sql.NullInt64   `db:"panel_count"`
	BatteryCount sql.NullInt64   `db:"battery_count"`
	Voltage      sql.NullFloat64 `db:"voltage"`
}

func (r *SolarPanelPostgres) AllData() ([]*SPAllData, error) {
	query := `
		SELECT 
			t1.nominal_power,
			t1.recorded_at,
			t2.length,
			t2.width,
			t2.weight,
			t2.panel_count,
			t2.battery_count,
			t2.voltage,
			t3.model_name,
			t4.brand_name
		FROM solar_panel_power_history t1
		INNER JOIN solar_panel t2 ON t1.panel_id = t2.id
		INNER JOIN model t3 ON t2.model_id = t3.id
		INNER JOIN brand t4 ON t3.brand_id = t4.id
	`

	var characteristics []*SPAllData
	if err := r.db.Select(&characteristics, query); err != nil {
		return nil, fmt.Errorf("failed to get solar panel data: %v", err)
	}

	return characteristics, nil
}
