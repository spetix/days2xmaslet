package dateutil

import (
	"fmt"
	"time"
)

const (
	Day = time.Duration(time.Hour * 24)
)

func GetUnitString(unit time.Duration) string {
	switch unit {
	case Day:
		return "days"
	case time.Hour:
		return "hours"
	case time.Minute:
		return "minutes"
	case time.Second:
		return "seconds"
	default:
		return "days"
	}
}

func StringToDateNextOccurrence(s string) time.Time {
	t, _ := time.Parse("01-02", s)
	t = time.Date(time.Now().Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	if t.Before(time.Now()) {
		t = t.AddDate(0, 0, 1)
	}
	return t
}

func Format(d time.Duration, unit time.Duration) string {

	switch unit {
	case Day:
		return fmt.Sprintf("%d", d.Truncate(Day)/Day)
	case time.Hour:
		{
			return fmt.Sprint(d.Truncate(time.Hour).Hours(), "h")
		}

	case time.Minute:
		return fmt.Sprint((d.Truncate(time.Minute) + time.Minute).String(), "m")
	case time.Second:
		return fmt.Sprint((d.Truncate(time.Second) + time.Second).String(), "s")
	default:
		return d.String()
	}
}

func GetUnit(unit string) time.Duration {
	switch unit {
	case "d":
		return Day
	case "h":
		return time.Hour
	case "m":
		return time.Minute
	case "s":
		return time.Second
	default:
		return Day
	}
}
