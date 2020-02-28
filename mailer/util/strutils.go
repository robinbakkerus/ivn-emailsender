package util

import (
	"strconv"
	"time"
)

// ToStr ..
func ToStr(nr int) string {
	return strconv.Itoa(nr)
}

// DateToStr ..
func DateToStr(date time.Time) string {
	return ToStr(date.Year()) + "-" + date.Month().String() + "-" + ToStr(date.Local().Day())
}
