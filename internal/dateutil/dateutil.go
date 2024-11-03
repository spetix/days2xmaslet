package dateutil

import (
	"fmt"
	"time"
)

const (
	Day = time.Duration(time.Hour * 24)
)

func HowManyDaysToXmas(actualDate time.Time, unit time.Duration) time.Duration {

	adjustedDate := actualDate.Truncate(unit)

	xmas := time.Date(adjustedDate.Year(), 12, 25, 0, 0, 0, 0, adjustedDate.Location())
	if adjustedDate.After(xmas) {
		xmas = xmas.AddDate(1, 0, 0)
	}
	return xmas.Sub(adjustedDate)
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
