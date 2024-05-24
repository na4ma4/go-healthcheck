package healthcheck_test

import (
	"errors"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	health "github.com/na4ma4/go-healthcheck"
)

func TestCoreWithCallbacks(t *testing.T) {
	t.Parallel()
	callbacks := map[health.Status]bool{}

	genCallback := func(statusConst health.Status) health.ItemCallbackFunc {
		return func(status health.Status, item health.Item, err error) bool {
			t.Logf("callback called status[%s] on item[%s]", status, item.Name())
			if item.Name() != "test2" {
				return true
			}
			if status != statusConst {
				t.Errorf(
					"coreWithCallback: callback status : got '%s' expected '%s'",
					status, statusConst,
				)
			}

			if status == health.StatusErrored {
				if !errors.Is(err, errFoobar) {
					t.Errorf("coreWithCallback: error to callback got '%s', expected '%s'", err, errFoobar)
				}
			}

			callbacks[status] = true

			return true
		}
	}

	corecb := health.NewCoreWithCallbacks(
		health.AddOnStatusCallback(health.StatusStarting, genCallback(health.StatusStarting)),
		health.AddOnStatusCallback(health.StatusRunning, genCallback(health.StatusRunning)),
		health.AddOnStatusCallback(health.StatusErrored, genCallback(health.StatusErrored)),
		health.AddOnStatusCallback(health.StatusFinished, genCallback(health.StatusFinished)),
	)

	corecb.Get("test1")
	test2 := corecb.Get("test2")
	corecb.Get("test1")

	expect := map[health.Status]bool{ //nolint:exhaustive // completing over time.
		health.StatusStarting: true,
	}

	if diff := cmp.Diff(callbacks, expect); diff != "" {
		t.Errorf("coreWithCallback: callbacks executed -got +want:\n%s", diff)
	}

	test2.Start()
	expect[health.StatusRunning] = true
	if diff := cmp.Diff(callbacks, expect); diff != "" {
		t.Errorf("coreWithCallback: callbacks executed -got +want:\n%s", diff)
	}

	test2.Error(errFoobar)
	expect[health.StatusErrored] = true
	if diff := cmp.Diff(callbacks, expect); diff != "" {
		t.Errorf("coreWithCallback: callbacks executed -got +want:\n%s", diff)
	}

	test2.Stop()
	expect[health.StatusFinished] = true
	if diff := cmp.Diff(callbacks, expect); diff != "" {
		t.Errorf("coreWithCallback: callbacks executed -got +want:\n%s", diff)
	}

	td1 := test2.Duration()
	time.Sleep(time.Millisecond)
	corecb.Stop("test2")
	td2 := test2.Duration()
	if td1 != td2 {
		t.Errorf(
			"coreWithCallback: item duration got '%s' and '%s' but "+
				"expected to be unchanged after calling stop again",
			td1.String(),
			td2.String(),
		)
	}
}

func TestCoreWithCallbacks_Stopped(t *testing.T) {
	t.Parallel()

	core := health.NewCoreWithCallbacks()

	core.Get("test1").Start()
	core.Get("test2").Start()
	core.Get("test3").Start()
	core.Stop("test2")

	expect := map[string]bool{
		"test1": true,
		"test2": true,
		"test3": true,
	}
	if diff := cmp.Diff(core.Status(), expect); diff != "" {
		t.Errorf("healthcheck.Status(): status -got +want:\n%s", diff)
	}

	expect = map[string]bool{
		"test1": true,
		"test2": false,
		"test3": true,
	}
	core.Get("test2").Start().Error(errors.New("foo"))

	if diff := cmp.Diff(core.Status(), expect); diff != "" {
		t.Errorf("healthcheck.Status(): status -got +want:\n%s", diff)
	}
}
