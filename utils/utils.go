package utils

import (
	"strings"
	"time"
)

func ParseStringToTime(s string) (date time.Time, err error) {
	date, err = time.Parse("2006-01-02", s)
	if err != nil {
		return
	}

	return
}

func CheckIfDuplicateKeyError(err error) bool {
	s := err.Error()
	return strings.Contains(s, "duplicate key value violates unique constraint")
}
