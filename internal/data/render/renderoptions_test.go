package render

import (
	"testing"
	"time"

	"github.com/spetix/days2xmasleft/internal/dateutil"
)

type dateAndExpectedDays struct {
	Date                    time.Time
	ExpectedDays            time.Duration
	Unit                    time.Duration
	ExpectedFormattedResult string
}

var currentDateAndExpected = []dateAndExpectedDays{
	{time.Date(2024, 12, 24, 0, 0, 0, 0, time.UTC), dateutil.Day, dateutil.Day, "1"},
	{time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC), 0, dateutil.Day, "0"},
	{time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC), 364 * dateutil.Day, dateutil.Day, "364"},
	{time.Date(2024, 12, 27, 0, 0, 0, 0, time.UTC), 363 * dateutil.Day, dateutil.Day, "363"},
	{time.Date(2019, 12, 28, 0, 0, 0, 0, time.UTC), 363 * dateutil.Day, dateutil.Day, "363"},
}

func TestGetUnit(t *testing.T) {

	for _, v := range currentDateAndExpected {
		r := RenderOptions{Unit: v.Unit}
		unit := r.FormatData(v.ExpectedDays)
		if unit != v.ExpectedFormattedResult {
			t.Fail()
		}
	}
}
