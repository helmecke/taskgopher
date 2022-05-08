package timeutils

import (
	"testing"
	"time"
)

func TestDiff(t *testing.T) {

	t.Run("seconds", func(t *testing.T) {
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := Diff(dt, dt.Add(time.Second))
		want := "1s"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("minutes", func(t *testing.T) {
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := Diff(dt, dt.Add(time.Minute))
		want := "1m"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("hours", func(t *testing.T) {
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := Diff(dt, dt.Add(time.Hour))
		want := "1h"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("days", func(t *testing.T) {
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := Diff(dt, dt.Add(time.Hour*24))
		want := "1d"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("months", func(t *testing.T) {
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := Diff(dt, dt.Add(time.Hour*24*7*5))
		want := "1M"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("years", func(t *testing.T) {
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := Diff(dt, dt.Add(time.Hour*24*7*53))
		want := "1y"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
