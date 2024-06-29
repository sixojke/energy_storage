package domain

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type Int4Range struct {
	Lower int
	Upper int
}

type NumRange struct {
	Lower float64
	Upper float64
}

// Scan - Implement the Scanner interface for Int4Range
func (r *Int4Range) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return r.scanFromString(string(v))
	case string:
		return r.scanFromString(v)
	default:
		return fmt.Errorf("could not convert %v to Int4Range", value)
	}
}

// Вспомогательный метод для обработки строк
func (r *Int4Range) scanFromString(str string) error {
	str = strings.TrimSpace(str)
	if len(str) < 2 {
		return fmt.Errorf("invalid Int4Range format: %v", str)
	}

	str = str[1 : len(str)-1] // Удаляем первую и последнюю скобки
	parts := strings.Split(str, ",")

	if len(parts) != 2 {
		return fmt.Errorf("invalid Int4Range format: %v", str)
	}

	lower, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return err
	}

	upperStr := strings.TrimSpace(parts[1])
	upperStr = strings.TrimRight(upperStr, ")")
	upperStr = strings.TrimRight(upperStr, "]")
	upper, err := strconv.Atoi(upperStr)
	if err != nil {
		return err
	}

	r.Lower = lower
	r.Upper = upper
	return nil
}

// Scan - Implement the Scanner interface for NumRange
func (r *NumRange) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return r.scanFromString(string(v))
	case string:
		return r.scanFromString(v)
	default:
		return fmt.Errorf("could not convert %v to NumRange", value)
	}
}

// Вспомогательный метод для обработки строк
func (r *NumRange) scanFromString(str string) error {
	str = strings.TrimSpace(str)
	if len(str) < 2 {
		return fmt.Errorf("invalid NumRange format: %v", str)
	}

	str = str[1 : len(str)-1] // Удаляем первую и последнюю скобки
	parts := strings.Split(str, ",")

	if len(parts) != 2 {
		return fmt.Errorf("invalid NumRange format: %v", str)
	}

	lower, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return err
	}

	upperStr := strings.TrimSpace(parts[1])
	upperStr = strings.TrimRight(upperStr, ")")
	upperStr = strings.TrimRight(upperStr, "]")
	upper, err := strconv.ParseFloat(upperStr, 64)
	if err != nil {
		return err
	}

	r.Lower = lower
	r.Upper = upper
	return nil
}

// Value - Implement the Valuer interface for Int4Range
func (r Int4Range) Value() (driver.Value, error) {
	return fmt.Sprintf("[%d,%d]", r.Lower, r.Upper), nil
}

// Value - Implement the Valuer interface for NumRange
func (r NumRange) Value() (driver.Value, error) {
	return fmt.Sprintf("[%f,%f]", r.Lower, r.Upper), nil
}
