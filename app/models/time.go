package models

import (
	"fmt"
	"time"
)

func GetTime() (year string, month string, day string) {
	t := time.Now()
	year = fmt.Sprintf("%d", t.Year())
	month = fmt.Sprintf("%02d", t.Month())
	day = fmt.Sprintf("%02d", t.Day())
	return year, month, day
}
