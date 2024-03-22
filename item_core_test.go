package healthcheck_test

import (
	"testing"
	"time"

	health "github.com/na4ma4/go-healthcheck"
)

func TestItemCoreSetStatus(t *testing.T) {
	item := health.NewItemCore("foo")

	if item.Status() != health.StatusStarting {
		t.Errorf("item.Status: get '%s', want '%s'",
			item.Status(), health.StatusStarting,
		)
	}

	item.Start()
	if item.Status() != health.StatusRunning {
		t.Errorf("item.Status: get '%s', want '%s'",
			item.Status(), health.StatusRunning,
		)
	}

	item.Error(errFoobar)
	if item.Status() != health.StatusErrored {
		t.Errorf("item.Status: get '%s', want '%s'",
			item.Status(), health.StatusErrored,
		)
	}

	item.SetStatus(health.StatusRunning)
	if item.Status() != health.StatusRunning {
		t.Errorf("item.Status: get '%s', want '%s'",
			item.Status(), health.StatusRunning,
		)
	}

	item.Stop()
	if item.Status() != health.StatusFinished {
		t.Errorf("item.Status: get '%s', want '%s'",
			item.Status(), health.StatusFinished,
		)
	}

	item.SetStatus(health.StatusRunning)
	if item.Status() != health.StatusRunning {
		t.Errorf("item.Status: get '%s', want '%s'",
			item.Status(), health.StatusRunning,
		)
	}

	item.Error(nil)
	if item.Status() != health.StatusFinished {
		t.Errorf("item.Status: get '%s', want '%s'",
			item.Status(), health.StatusFinished,
		)
	}
}

func TestItemCoreErrorIsNil(t *testing.T) {
	item := health.NewItemCore("foo").Start()

	if item.Status() != health.StatusRunning {
		t.Errorf("item.Status: get '%s', want '%s'",
			item.Status(), health.StatusRunning,
		)
	}

	item.Error(nil)
	if item.Status() != health.StatusFinished {
		t.Errorf("item.Status: get '%s', want '%s'",
			item.Status(), health.StatusFinished,
		)
	}
}

func TestItemCoreStartTime(t *testing.T) {
	item := health.NewItemCore("foo")
	item.Start()

	ts1 := item.StartTime()
	time.Sleep(time.Millisecond)
	ts2 := item.StartTime()

	if ts1 != ts2 {
		t.Errorf("item.StartTime: got '%s', want '%s'", ts1.String(), ts2.String())
	}

	item.SetStatus(health.StatusStarting)
	ts3 := item.StartTime()

	if ts2 == ts3 {
		t.Errorf("item.StartTime: got '%s', want changed value", ts3.String())
	}
}
