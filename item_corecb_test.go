package healthcheck_test

import (
	"testing"
	"time"

	health "github.com/na4ma4/go-healthcheck"
)

func TestItemCoreWithCallbacks(t *testing.T) {
	t.Parallel()

	startTime := time.Now()
	item := health.NewItemCoreWithCallbacks("test1")
	time.Sleep(time.Millisecond)
	item.Start()
	startedTime := time.Now()

	time.Sleep(time.Millisecond)
	item.Stop()

	lifecycle := item.Lifecycle()

	if len(lifecycle) != 3 {
		t.Fatalf("item.Lifecycle: lifecycle items got '%d' want '%d'", len(lifecycle), 3)
	}

	if v := lifecycle[1].Timestamp.Sub(lifecycle[0].Timestamp.Time); v < time.Millisecond {
		t.Errorf("item.Lifecycle: got '%s', expect difference between 0 and 1 to be more than 1us", v.String())
	}

	if v := lifecycle[2].Timestamp.Sub(lifecycle[1].Timestamp.Time); v < time.Millisecond {
		t.Errorf("item.Lifecycle: got '%s', expect difference between 1 and 2 to be more than 1us", v.String())
	}

	if item.Name() != "test1" {
		t.Errorf("item.Name: got '%s' want '%s'", item.Name(), "test1")
	}

	if !startedTime.After(item.StartTime().Time) || startTime.After(item.StartTime().Time) {
		t.Error("item.StartTime: not in between timestamps before and after starting")
	}
}

// TestItemCoreWithCallbacks_AbortAll tests to make sure that no events will be stored if
// the callbacks disable all of them.
func TestItemCoreWithCallbacks_AbortAll(t *testing.T) {
	t.Parallel()

	genCallback := func() health.ItemCallbackFunc {
		return func(status health.Status, item health.Item, _ error) bool {
			t.Logf("callback called status[%s] on item[%s]", status, item.Name())
			return false
		}
	}
	item := health.NewItemCoreWithCallbacks(
		"test1",
		health.AddOnStatusCallback(health.StatusStarting, genCallback()),
		health.AddOnStatusCallback(health.StatusRunning, genCallback()),
		health.AddOnStatusCallback(health.StatusErrored, genCallback()),
		health.AddOnStatusCallback(health.StatusFinished, genCallback()),
	)
	time.Sleep(time.Millisecond)
	item.Start()
	time.Sleep(time.Millisecond)
	item.Stop()
	time.Sleep(time.Millisecond)
	item.Error(errFoobar)

	lifecycle := item.Lifecycle()

	if len(lifecycle) != 0 {
		t.Fatalf("item.Lifecycle: lifecycle items got '%d' want '%d'", len(lifecycle), 0)
	}
}

func TestItemCoreWithCallbacks_AbortAll_CallingStart(t *testing.T) {
	t.Parallel()

	genCallback := func() health.ItemCallbackFunc {
		return func(status health.Status, item health.Item, _ error) bool {
			t.Logf("callback called status[%s] on item[%s]", status, item.Name())
			if status == health.StatusRunning {
				t.Logf("calling Start() for status[%s] sub-item[%s]", status, item.Name())
				item.Start()
			}
			return false
		}
	}
	item := health.NewItemCoreWithCallbacks(
		"test1",
		health.AddOnStatusCallback(health.StatusStarting, genCallback()),
		health.AddOnStatusCallback(health.StatusRunning, genCallback()),
		health.AddOnStatusCallback(health.StatusErrored, genCallback()),
		health.AddOnStatusCallback(health.StatusFinished, genCallback()),
	)
	time.Sleep(time.Millisecond)
	item.Start()
	time.Sleep(time.Millisecond)
	item.Stop()
	time.Sleep(time.Millisecond)
	item.Error(errFoobar)

	lifecycle := item.Lifecycle()

	if len(lifecycle) != 1 {
		t.Fatalf("item.Lifecycle: lifecycle items got '%d' want '%d'", len(lifecycle), 1)
	}
}
