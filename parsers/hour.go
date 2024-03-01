package parsers

import (
	"time"
)

const layoutTime = "15:04"

func ParseHour(t string) (time.Duration, error) {
	tm, err := time.Parse(layoutTime, t)
	if err != nil {
		return 0, err
	}

	duration := (time.Duration(tm.Hour()) * time.Hour) + (time.Duration(tm.Minute()) * time.Minute)
	return duration, nil
}
