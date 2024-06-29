package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type WindTurbinePostgres struct {
	db *sqlx.DB
}

func NewWindTurbinePostgres(db *sqlx.DB) *WindTurbinePostgres {
	return &WindTurbinePostgres{db: db}
}

type WTAllData struct {
	Power      sql.NullFloat64 `db:"power"`
	Voltage    sql.NullFloat64 `db:"voltage"`
	RecordedAt time.Time       `db:"recorded_at"`
	Model      sql.NullString  `db:"model_name"`
	Brand      sql.NullString  `db:"brand_name"`
}

func (r *WindTurbinePostgres) AllData() ([]*WTAllData, error) {
	query := `
		SELECT
			t1.power,
			t1.voltage,
			t1.recorded_at,
			t2.model_name,
			t3.brand_name
		FROM wind_turbine_history t1
		INNER JOIN model t2 ON t1.model_id = t2.id
		INNER JOIN brand t3 ON t2.brand_id = t3.id
	`

	var characteristics []*WTAllData
	if err := r.db.Select(&characteristics, query); err != nil {
		return nil, fmt.Errorf("failed to get wind turbine data: %v", err)
	}

	return characteristics, nil
}
