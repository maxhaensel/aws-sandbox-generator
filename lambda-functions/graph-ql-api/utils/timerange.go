package utils

import (
	"time"
)

func TimeRange(year int, month int, day int) (*string, *string) {
	a := time.Now().UTC().Format(time.RFC3339)
	b := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Now().Location()).UTC().Format(time.RFC3339)
	return &a, &b
}
