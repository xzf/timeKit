package timeKit

import "time"

const (
	Day   = time.Hour * 24
	Week  = Day * 7
	Month = Day * 30  //not always a standard month, it may be wrong
	Year  = Day * 365 //not always a standard year, it may be wrong
)

func ToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func GetDateEnd(t time.Time) time.Time {
	t = ToDate(t)
	return t.Add(time.Nanosecond * -1)
}

func GetMonthStart(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func GetMonthEnd(t time.Time) time.Time {
	t = time.Date(t.Year(), t.Month(), 28, 0, 0, 0, 0, t.Location())
	oldMonth := t.Month()
	for oldMonth == t.Month() {
		t = t.Add(Day)
	}
	return t.Add(Day * -1)
}

func GetYearStart(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

func GetYearEnd(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 0, 0, 0, 0, t.Location())
}

func GetLastDay(t time.Time) time.Time {
	return t.Add(Day * -1)
}

func GetNextDay(t time.Time) time.Time {
	return t.Add(Day)
}
