package healthcheck_test

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	health "github.com/na4ma4/go-healthcheck"
)

func TestHealthEventStringer(t *testing.T) {
	timeString := "2006-01-02T15:04:05Z"
	ts, _ := time.Parse(time.RFC3339, timeString)
	event := health.Event{health.NewEventTime(ts), health.NewEventStatus(health.StatusStarting)}

	expect := "STARTING[" + timeString + "]"

	if diff := cmp.Diff(event.String(), expect); diff != "" {
		t.Errorf("health.Event.String: callbacks executed -got +want:\n%s", diff)
	}
}

type jsonTest struct {
	Status health.EventStatus `json:"status"`
}

func TestUnmarshalMarshalEventStatus(t *testing.T) {
	test := jsonTest{
		Status: health.NewEventStatus(health.Status_RUNNING),
	}

	buf := bytes.NewBuffer(nil)

	if err := json.NewEncoder(buf).Encode(test); err != nil {
		t.Errorf("unexpected error encoding EventStatus to json: %s", err)
		t.FailNow()
	}

	testOut := jsonTest{}

	t.Logf("JSON Encoded Status: %s", buf.Bytes())

	if err := json.NewDecoder(buf).Decode(&testOut); err != nil {
		t.Errorf("unexpected error decoding EventStatus from json: %s", err)
		t.FailNow()
	}

	if test.Status != testOut.Status {
		t.Errorf("test output does not equal test input: got '%s', want '%s'", testOut.Status, test.Status)
	}
}
