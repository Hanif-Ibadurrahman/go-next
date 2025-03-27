package util

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}
type IntOrString int
type FloatOrString float64

const timeFormat = "2006-01-02 15:04:05"

// ✅ JSON Unmarshal function to parse different time formats
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	// Handle null or empty values
	if str == "null" || str == "" {
		ct.Time = time.Time{}
		return nil
	}

	// Define possible formats
	formats := []string{
		"2006-01-02 15:04:05",       // Full datetime
		"2006-01-02T15:04:05Z07:00", // ISO 8601
		"2006-01-02",                // Date-only
	}

	var err error
	for _, format := range formats {
		ct.Time, err = time.Parse(format, str)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("invalid time format: %s", str)
}

// ✅ Make `CustomTime` compatible with MySQL
func (ct CustomTime) Value() (driver.Value, error) {
	if ct.IsZero() {
		return nil, nil // Return NULL for empty timestamps
	}
	return ct.Format(timeFormat), nil
}

// ✅ Convert MySQL DATETIME result to `CustomTime`
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		*ct = CustomTime{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*ct = CustomTime{Time: v}
	case string:
		parsedTime, err := time.Parse(timeFormat, v)
		if err != nil {
			return fmt.Errorf("failed to parse time: %w", err)
		}
		*ct = CustomTime{Time: parsedTime}
	default:
		return fmt.Errorf("unsupported type for CustomTime: %T", value)
	}

	return nil
}

func (i *IntOrString) UnmarshalJSON(b []byte) error {
	var n int
	if err := json.Unmarshal(b, &n); err == nil {
		*i = IntOrString(n)
		return nil
	}

	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		num, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("invalid noindex format: %s", s)
		}
		*i = IntOrString(num)
		return nil
	}

	return fmt.Errorf("noindex must be int or string")
}

func (f *FloatOrString) UnmarshalJSON(b []byte) error {
	var num float64
	if err := json.Unmarshal(b, &num); err == nil {
		*f = FloatOrString(num)
		return nil
	}

	var str string
	if err := json.Unmarshal(b, &str); err == nil {
		parsed, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return fmt.Errorf("invalid float format: %s", str)
		}
		*f = FloatOrString(parsed)
		return nil
	}

	return fmt.Errorf("expected float or string")
}
