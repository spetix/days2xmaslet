package dateutil

import (
	"testing"
	"time"
)

type dateAndExpectedDay struct {
	Date                    time.Time
	ExpectedDays            time.Duration
	Unit                    time.Duration
	ExpectedFormattedResult string
}

var currentDateAndExpectedDays = []dateAndExpectedDay{
	{time.Date(2024, 12, 24, 0, 0, 0, 0, time.UTC), Day, Day, "1"},
	{time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC), 0, Day, "0"},
	{time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC), 364 * Day, Day, "364"},
	{time.Date(2024, 12, 27, 0, 0, 0, 0, time.UTC), 363 * Day, Day, "363"},
	{time.Date(2019, 12, 28, 0, 0, 0, 0, time.UTC), 363 * Day, Day, "363"},
}

func TestGetUnit(t *testing.T) {
	unit := GetUnit("d")
	if unit != Day {
		t.Fail()
	}
	unit = GetUnit("h")
	if unit != time.Hour {
		t.Fail()
	}
	unit = GetUnit("m")
	if unit != time.Minute {
		t.Fail()
	}
	unit = GetUnit("s")
	if unit != time.Second {
		t.Fail()
	}
}
