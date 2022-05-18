package optimizer

import (
	"math"
	"time"
)

func RoundNumber(number float64, decimals int) float64 {
	factor := math.Pow10(decimals)
	globalCounter++
	return math.Round(number*factor) / factor
}

func dateToYearMonth(day string) (string, time.Time) {
	layout1 := "2006-01-02"
	layout2 := "2006-01"
	tmpDay, _ := time.Parse(layout1, day)

	return tmpDay.Format(layout2), tmpDay
}

func isDateBetween(day time.Time, start time.Time, end time.Time) bool {
	return (day.After(start) && day.Before(end)) || day.Equal(start) || day.Equal(end)
}
