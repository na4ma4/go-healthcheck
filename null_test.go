package healthcheck_test

import (
	"testing"

	health "github.com/na4ma4/go-healthcheck"
)

func TestNullCoverage(t *testing.T) {
	core := health.NewNull()
	core.Status()
	core.Stop("test2")
	testNullCoverage(t, core)
}

func testNullCoverage(t *testing.T, core health.Health) {
	item := core.Get("test1").Start().Stop().Error(nil)

	if v := item.Status(); v != health.StatusUnknown { // null doesn't do changes, always returns Unknown
		t.Errorf("nullItem.Status: got '%s' want '%s'", v, health.StatusUnknown)
	}

	if v := item.Name(); v != "null" {
		t.Errorf("nullItem.Name: got '%s' want 'null'", v)
	}

	if v := item.Duration(); v.Seconds() != 0 {
		t.Errorf("nullItem.Duration: got '%s' want '0s'", v.String())
	}

	if v := item.StartTime(); !v.IsZero() {
		t.Errorf("nullItem.StartTime: got '%s' want 'time.Time{}'", v.String())
	}

	if v := item.Lifecycle(); len(v) > 0 {
		t.Errorf("nullItem.Lifecycle: returned '%d' items want '0'", len(v))
	}
}
