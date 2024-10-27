package dateutil

import (
	"fmt"
	"time"
)

const (
	Day = time.Duration(time.Hour * 24)
)

func HowManyDaysToXmas(actualDate time.Time) time.Duration {

	xmas := time.Date(actualDate.Year(), 12, 25, 0, 0, 0, 0, actualDate.Location())
	if actualDate.After(xmas) {
		xmas = xmas.AddDate(1, 0, 0)
	}
	return xmas.Sub(actualDate)
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
		return d.Truncate(time.Minute).String()
	case time.Second:
		return d.Truncate(time.Second).String()
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
