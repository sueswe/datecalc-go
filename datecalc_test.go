package main

import (
	"testing"
	"time"
)

func TestFirstFriday(t *testing.T) {
	date := "20251101"
	want := "2025-11-07"

	n := nth_dow_of_month(date, 1, "Friday")
	if n != want {
		t.Errorf("Result does not fit: %s, want: %s", n, want)
	}

}

func TestAddDateDay(t *testing.T) {
	date := "20251101"
	want := "Wednesday,2025-11-05"

	n := addSubtr(date, 4, 0, 0)
	if n != want {
		t.Errorf("Result does not fit: %s, want: %s", n, want)
	}
}

func TestAddDateMonth(t *testing.T) {
	date := "20251106"
	want := "Tuesday,2026-01-06"

	n := addSubtr(date, 0, 2, 0)
	if n != want {
		t.Errorf("Result does not fit: %s, want: %s", n, want)
	}
}

func TestAddDateYear(t *testing.T) {
	date := "20251106"
	want := "Friday,2026-11-06"

	n := addSubtr(date, 0, 0, 1)
	if n != want {
		t.Errorf("Result does not fit: %s , want: %s", n, want)
	}
}

func TestDaysInMonth(t *testing.T) {
	dateString := "20251003"
	format := "20060102"
	c, _ := time.Parse(format, dateString)
	n := daysInMonth(c)
	want := 31
	if n != want {
		t.Errorf("Result does not fit: %d , want: %d", n, want)
	}
}
