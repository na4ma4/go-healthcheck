package healthcheck_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/na4ma4/go-healthcheck"
)

func getReport[T healthcheck.Report](ts *httptest.Server, queryString string) (T, error) {
	data := new(T)
	var resp *http.Response
	{
		var err error
		resp, err = http.Get(ts.URL + "?" + queryString)
		if err != nil {
			return *data, err
		}
	}

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return *data, err
	}

	return *data, nil
}

func TestHandler_GlobalStatus(t *testing.T) {
	tests := []struct {
		description string
		item        func(healthcheck.Health)
		expect      healthcheck.Status
	}{
		{
			"starting = starting",
			func(c healthcheck.Health) {
				c.Get("Starting")
			},
			healthcheck.StatusStarting,
		},
		{
			"running+starting = starting",
			func(c healthcheck.Health) {
				c.Get("Starting")
				c.Get("Running").Start()
			},
			healthcheck.StatusStarting,
		},
		{
			"running+finished = running",
			func(c healthcheck.Health) {
				c.Get("Running").Start()
				c.Get("Finished").Start().Stop()
			},
			healthcheck.StatusRunning,
		},
		{
			"running+finished+errored = errored",
			func(c healthcheck.Health) {
				c.Get("Running").Start()
				c.Get("Finished").Start().Stop()
				c.Get("Errored").Start().Error(errors.New("foo"))
			},
			healthcheck.StatusErrored,
		},
		{
			"running+finished+errored(nil) = running",
			func(c healthcheck.Health) {
				c.Get("Running").Start()
				c.Get("Finished").Start().Stop()
				c.Get("Errored").Start().Error(nil)
			},
			healthcheck.StatusRunning,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			hc := healthcheck.NewCore()
			tt.item(hc)
			ts := httptest.NewServer(healthcheck.Handler(hc))

			resp, err := getReport[*healthcheck.ReportStandard](ts, "")
			if err != nil {
				t.Errorf("unable to request or decode report: got error '%s'", err)
				return
			}

			if resp.Status.String() != tt.expect.String() {
				t.Errorf("got '%s', want '%s'", resp.Status.String(), tt.expect.String())
			}
		})
	}
}

func TestHandler_GlobalStatusSimple(t *testing.T) {
	tests := []struct {
		description string
		item        func(healthcheck.Health)
		expect      healthcheck.ReportStatus
	}{
		{
			"starting = starting",
			func(c healthcheck.Health) {
				c.Get("Starting")
			},
			healthcheck.ReportStatusYellow,
		},
		{
			"running+starting = starting",
			func(c healthcheck.Health) {
				c.Get("Starting")
				c.Get("Running").Start()
			},
			healthcheck.ReportStatusYellow,
		},
		{
			"running+finished = running",
			func(c healthcheck.Health) {
				c.Get("Running").Start()
				c.Get("Finished").Start().Stop()
			},
			healthcheck.ReportStatusGreen,
		},
		{
			"running+finished+errored(no start) = errored",
			func(c healthcheck.Health) {
				c.Get("Running").Start()
				c.Get("Finished").Start().Stop()
				c.Get("Errored").Error(errors.New("foo"))
			},
			healthcheck.ReportStatusRed,
		},
		{
			"running+finished+errored(started) = errored",
			func(c healthcheck.Health) {
				c.Get("Running").Start()
				c.Get("Finished").Start().Stop()
				c.Get("Errored").Start().Error(errors.New("foo"))
			},
			healthcheck.ReportStatusRed,
		},
		{
			"running+finished+errored(nil) = running",
			func(c healthcheck.Health) {
				c.Get("Running").Start()
				c.Get("Finished").Start().Stop()
				c.Get("Errored").Start().Error(nil)
			},
			healthcheck.ReportStatusGreen,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			hc := healthcheck.NewCore()
			tt.item(hc)
			ts := httptest.NewServer(healthcheck.Handler(hc))

			resp, err := getReport[*healthcheck.ReportSimple](ts, "simple=1")
			if err != nil {
				t.Errorf("unable to request or decode report: got error '%s'", err)
				return
			}

			if resp.Status.String() != tt.expect.String() {
				t.Errorf("got '%s', want '%s'", resp.Status.String(), tt.expect.String())
				t.Logf("Response: %+v", resp)
			}
		})
	}
}

func TestHandler_SimpleStatus(t *testing.T) {
	hc := healthcheck.NewCore()

	cmpoptions := []cmp.Option{
		cmpopts.IgnoreUnexported(healthcheck.ReportSimple{}),
		cmpopts.IgnoreUnexported(healthcheck.ReportItemSimple{}),
	}

	successCheck1 := hc.Get("Success1")
	successCheck2 := hc.Get("Success2").Start()
	ts := httptest.NewServer(healthcheck.Handler(hc))

	// Test Starting and Started
	resp, err := getReport[*healthcheck.ReportSimple](ts, "simple=1")
	if err != nil {
		t.Errorf("unable to request or decode report: got error '%s'", err)
		return
	}

	expect := &healthcheck.ReportSimple{
		Status: healthcheck.ReportStatusText(healthcheck.ReportStatusYellow.String()),
		Services: []*healthcheck.ReportItemSimple{
			{Name: "Success1", Status: healthcheck.ReportStatusText(healthcheck.ReportStatusYellow.String())},
			{Name: "Success2", Status: healthcheck.ReportStatusText(healthcheck.ReportStatusGreen.String())},
		},
	}
	if diff := cmp.Diff(resp, expect, cmpoptions...); diff != "" {
		t.Errorf("Starting+Started: -got +want:\n%s", diff)
	}

	// Test All Started
	successCheck1.Start()
	expect.Status = healthcheck.ReportStatusText(healthcheck.ReportStatusGreen.String())
	expect.Services[0].Status = healthcheck.ReportStatusText(healthcheck.ReportStatusGreen.String())

	resp, err = getReport[*healthcheck.ReportSimple](ts, "simple=1")
	if err != nil {
		t.Errorf("unable to request or decode report: got error '%s'", err)
		return
	}
	if diff := cmp.Diff(resp, expect, cmpoptions...); diff != "" {
		t.Errorf("Started+Started: -got +want:\n%s", diff)
	}

	// Test Finished
	successCheck1.Stop()

	resp, err = getReport[*healthcheck.ReportSimple](ts, "simple=1")
	if err != nil {
		t.Errorf("unable to request or decode report: got error '%s'", err)
		return
	}

	if diff := cmp.Diff(resp, expect, cmpoptions...); diff != "" {
		t.Errorf("Stopped+Running: -got +want:\n%s", diff)
	}

	// Test Error
	successCheck2.Error(errors.New("foo"))
	expect.Status = healthcheck.ReportStatusText(healthcheck.ReportStatusRed.String())
	expect.Services[1].Status = healthcheck.ReportStatusText(healthcheck.ReportStatusRed.String())

	resp, err = getReport[*healthcheck.ReportSimple](ts, "simple=1")
	if err != nil {
		t.Errorf("unable to request or decode report: got error '%s'", err)
		return
	}

	if diff := cmp.Diff(resp, expect, cmpoptions...); diff != "" {
		t.Errorf("Stopped+Errored Success: -got +want:\n%s", diff)
	}
}
