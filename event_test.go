package healthcheck_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	health "github.com/na4ma4/go-healthcheck"
)

func TestHealthEventStringer(t *testing.T) {
	timeString := "2006-01-02T15:04:05Z"
	ts, _ := time.Parse(time.RFC3339, timeString)
	event := health.Event{health.NewEventTime(ts), health.StatusStarting}

	expect := "STARTING[" + timeString + "]"

	if diff := cmp.Diff(event.String(), expect); diff != "" {
		t.Errorf("health.Event.String: callbacks executed -got +want:\n%s", diff)
	}
}
