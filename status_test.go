package healthcheck_test

import (
	"fmt"
	"testing"

	"github.com/na4ma4/go-healthcheck"
)

func TestStatusLess(t *testing.T) {
	tests := []struct {
		x, y     healthcheck.Status
		lessThan bool
	}{
		// // Unknown(0) = Unknown
		// {0, healthcheck.StatusStarting, false},
		// {0, healthcheck.StatusRunning, false},
		// {0, healthcheck.StatusFinished, false},
		// {0, healthcheck.StatusUnknown, false},
		// {0, healthcheck.StatusErrored, true},

		// Errored = 0
		{healthcheck.StatusErrored, healthcheck.StatusErrored, false},
		{healthcheck.StatusErrored, healthcheck.StatusUnknown, false},
		{healthcheck.StatusErrored, healthcheck.StatusStarting, false},
		{healthcheck.StatusErrored, healthcheck.StatusRunning, false},
		{healthcheck.StatusErrored, healthcheck.StatusFinished, false},

		// Unknown = 1
		{healthcheck.StatusUnknown, healthcheck.StatusErrored, true},
		{healthcheck.StatusUnknown, healthcheck.StatusUnknown, false},
		{healthcheck.StatusUnknown, healthcheck.StatusStarting, false},
		{healthcheck.StatusUnknown, healthcheck.StatusRunning, false},
		{healthcheck.StatusUnknown, healthcheck.StatusFinished, false},

		// Starting = 2
		{healthcheck.StatusStarting, healthcheck.StatusErrored, true},
		{healthcheck.StatusStarting, healthcheck.StatusUnknown, true},
		{healthcheck.StatusStarting, healthcheck.StatusStarting, false},
		{healthcheck.StatusStarting, healthcheck.StatusRunning, false},
		{healthcheck.StatusStarting, healthcheck.StatusFinished, false},

		// Running = 3
		{healthcheck.StatusRunning, healthcheck.StatusErrored, true},
		{healthcheck.StatusRunning, healthcheck.StatusUnknown, true},
		{healthcheck.StatusRunning, healthcheck.StatusStarting, true},
		{healthcheck.StatusRunning, healthcheck.StatusRunning, false},
		{healthcheck.StatusRunning, healthcheck.StatusFinished, false},

		// Finished = 4
		{healthcheck.StatusFinished, healthcheck.StatusErrored, true},
		{healthcheck.StatusFinished, healthcheck.StatusUnknown, true},
		{healthcheck.StatusFinished, healthcheck.StatusStarting, true},
		{healthcheck.StatusFinished, healthcheck.StatusRunning, true},
		{healthcheck.StatusFinished, healthcheck.StatusFinished, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_lessthan_%s", tt.x, tt.y), func(t *testing.T) {
			if v := tt.x.Less(tt.y); v != tt.lessThan {
				t.Errorf("got '%t', want '%t'", v, tt.lessThan)
			}
		})
	}
}
