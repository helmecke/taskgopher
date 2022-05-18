package timeutils_test

import (
	"testing"
	"time"

	"github.com/helmecke/taskgopher/pkg/timeutils"
)

func TestDiff(t *testing.T) {
	t.Parallel()

	t.Run("seconds", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt, dt.Add(time.Second), false)
		want := "1s"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("seconds past", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt.Add(time.Second), dt, true)
		want := "-1s"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("seconds future", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt, dt.Add(time.Second), true)
		want := "+1s"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("minutes", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt, dt.Add(time.Minute), false)
		want := "1m"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("minutes past", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt.Add(time.Minute), dt, true)
		want := "-1m"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("minutes future", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt, dt.Add(time.Minute), true)
		want := "+1m"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("hours", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt, dt.Add(time.Hour), false)
		want := "1h"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("hours past", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt.Add(time.Hour), dt, true)
		want := "-1h"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("hours future", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt, dt.Add(time.Hour), true)
		want := "+1h"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("days", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt, dt.Add(time.Hour*24), false)
		want := "1d"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("days past", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt.Add(time.Hour*24), dt, true)
		want := "-1d"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("days future", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt, dt.Add(time.Hour*24), true)
		want := "+1d"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("months", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt, dt.Add(time.Hour*24*7*5), false)
		want := "1M"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("months past", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt.Add(time.Hour*24*7*5), dt, true)
		want := "-1M"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("months future", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt, dt.Add(time.Hour*24*7*5), true)
		want := "+1M"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("years", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt, dt.Add(time.Hour*24*7*53), false)
		want := "1y"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("years past", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt.Add(time.Hour*24*7*53), dt, true)
		want := "-1y"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("years future", func(t *testing.T) {
		t.Parallel()
		dt := time.Date(2020, 10, 32, 0, 0, 0, 0, time.UTC)

		got := timeutils.Diff(dt, dt.Add(time.Hour*24*7*53), true)
		want := "+1y"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
