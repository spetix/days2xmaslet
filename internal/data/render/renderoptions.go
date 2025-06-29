package render

import (
	"fmt"
	"time"

	"github.com/spetix/days2xmasleft/internal/dateutil"
)

type RenderOptions struct {
	Label           string
	Format          string
	ForegroundColor string
	BackgroundColor string
	Unit            time.Duration
}

func GetUnit(unit string) time.Duration {
	switch unit {
	case "d":
		return dateutil.Day
	case "h":
		return time.Hour
	case "m":
		return time.Minute
	case "s":
		return time.Second
	default:
		return dateutil.Day
	}
}

func (ro RenderOptions) FormatData(d time.Duration) string {

	switch ro.Unit {
	case dateutil.Day:
		return fmt.Sprintf("%d", d.Truncate(dateutil.Day)/dateutil.Day)
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
