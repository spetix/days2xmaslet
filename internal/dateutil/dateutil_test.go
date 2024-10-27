package dateutil

import (
	"testing"
	"time"
)

func TestHowManyDaysToXmas(t *testing.T) {
	actualDate := time.Date(2024, 12, 24, 0, 0, 0, 0, time.UTC)

	days := HowManyDaysToXmas(actualDate)
	if days != Day {
		t.Fail()
	}
}

func TestHowManyDaysToXmas_next(t *testing.T) {
	actualDate := time.Date(2024, 12, 27, 0, 0, 0, 0, time.UTC)

	days := HowManyDaysToXmas(actualDate)
	if days != Day*363 {
		t.Fail()
	}
}

func TestHowManyDaysToXmas_xmas(t *testing.T) {
	actualDate := time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)

	days := HowManyDaysToXmas(actualDate)
	if days != 0 {
		t.Fail()
	}
}

func TestHowManyDaysToXmas_leap(t *testing.T) {
	actualDate := time.Date(2019, 12, 27, 0, 0, 0, 0, time.UTC)

	days := HowManyDaysToXmas(actualDate)
	if days != Day*364 {
		t.Fail()
	}
}

func TestFormat_day(t *testing.T) {
	actualDate := time.Date(2024, 12, 24, 0, 0, 0, 0, time.UTC)

	formatted := Format(HowManyDaysToXmas(actualDate), Day)
	if formatted != "1" {
		t.Error("Exprected 1, got ", formatted)
	}
}

func TestFormat_day_next(t *testing.T) {
	actualDate := time.Date(2024, 12, 27, 0, 0, 0, 0, time.UTC)

	formatted := Format(HowManyDaysToXmas(actualDate), Day)
	if formatted != "363" {
		t.Error("Exprected 1, got ", formatted)
	}
}

func TestFormat(t *testing.T) {
	actualDate := time.Date(2024, 12, 24, 1, 0, 0, 0, time.UTC)

	formatted := Format(HowManyDaysToXmas(actualDate), time.Hour)
	if formatted != "23h" {
		t.Fail()
	}
}
