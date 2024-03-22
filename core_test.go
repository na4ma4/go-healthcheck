package healthcheck_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	health "github.com/na4ma4/go-healthcheck"
)

var errFoobar = errors.New("foobar")

func TestCoreStartStop(t *testing.T) {
	t.Parallel()

	core := health.NewCore()

	core.Get("test1").Start()
	core.Get("test2").Start()
	core.Get("test3").Start()
	core.Stop("test2")

	out := core.Status()
	expect := map[string]bool{
		"test1": true,
		"test2": false,
		"test3": true,
	}

	if diff := cmp.Diff(out, expect); diff != "" {
		t.Errorf("ReadZoneFile: status -got +want:\n%s", diff)
	}
}

func TestCoreStartErrored(t *testing.T) {
	t.Parallel()

	core := health.NewCore()

	core.Get("test1").Start()
	core.Get("test3").Start()
	test2 := core.Get("test2").Start()
	test2.Error(errFoobar)

	out := core.Status()
	expect := map[string]bool{
		"test1": true,
		"test2": false,
		"test3": true,
	}

	if diff := cmp.Diff(out, expect); diff != "" {
		t.Errorf("ReadZoneFile: status -got +want:\n%s", diff)
	}
}

func TestCoreStart_DurationChanged(t *testing.T) {
	t.Parallel()

	core := health.NewCore()
	test1 := core.Get("test1").Start()

	td1 := test1.Duration()
	time.Sleep(time.Millisecond)
	td2 := test1.Duration()

	if td1 == td2 {
		t.Errorf(
			"item.Duration: got '%s' for both but expected to be changing during execution",
			td1.String(),
		)
	}
}
func TestCoreStartErrored_DurationUnchanged(t *testing.T) {
	t.Parallel()

	core := health.NewCore()
	test1 := core.Get("test1").Start()
	test1.Error(errFoobar)

	td1 := test1.Duration()
	time.Sleep(time.Millisecond)
	td2 := test1.Duration()

	if td1 != td2 {
		t.Errorf(
			"item.Duration: got '%s' and '%s' but expected to be unchanged after error",
			td1.String(),
			td2.String(),
		)
	}
}

func TestCoreStartStopped_DurationUnchanged(t *testing.T) {
	t.Parallel()

	core := health.NewCore()
	test1 := core.Get("test1").Start()
	test1.Stop()

	td1 := test1.Duration()
	time.Sleep(time.Millisecond)
	td2 := test1.Duration()

	if td1 != td2 {
		t.Errorf(
			"item.Duration: got '%s' and '%s' but expected to be unchanged after stop",
			td1.String(),
			td2.String(),
		)
	}
}

func TestCoreMultipleStop_FirstDuration(t *testing.T) {
	t.Parallel()

	core := health.NewCore()
	test1 := core.Get("test1").Start()
	test1.Stop()
	td1 := test1.Duration()
	time.Sleep(time.Millisecond)

	core.Stop("test1")
	td2 := test1.Duration()

	if td1 != td2 {
		t.Errorf(
			"item.Duration: got '%s' and '%s' but expected to be unchanged after calling stop again",
			td1.String(),
			td2.String(),
		)
	}
}
