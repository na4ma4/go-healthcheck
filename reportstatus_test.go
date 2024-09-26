package healthcheck_test

import (
	"fmt"
	"testing"

	"github.com/na4ma4/go-healthcheck"
)

func TestReportStatusLess(t *testing.T) {
	tests := []struct {
		x, y     healthcheck.ReportStatus
		lessThan bool
	}{
		// Unknown(0) = NOTSET
		{0, healthcheck.ReportStatusRed, true},
		{0, healthcheck.ReportStatusYellow, true},
		{0, healthcheck.ReportStatusGreen, true},

		// Red = 0
		{healthcheck.ReportStatusRed, healthcheck.ReportStatusRed, false},
		{healthcheck.ReportStatusRed, healthcheck.ReportStatusYellow, false},
		{healthcheck.ReportStatusRed, healthcheck.ReportStatusGreen, false},

		// Yellow = 1
		{healthcheck.ReportStatusYellow, healthcheck.ReportStatusRed, true},
		{healthcheck.ReportStatusYellow, healthcheck.ReportStatusYellow, false},
		{healthcheck.ReportStatusYellow, healthcheck.ReportStatusGreen, false},

		// Green = 2
		{healthcheck.ReportStatusGreen, healthcheck.ReportStatusRed, true},
		{healthcheck.ReportStatusGreen, healthcheck.ReportStatusYellow, true},
		{healthcheck.ReportStatusGreen, healthcheck.ReportStatusGreen, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_lessthan_%s", tt.x, tt.y), func(t *testing.T) {
			if v := tt.x.Less(tt.y); v != tt.lessThan {
				t.Errorf("got '%t', want '%t'", v, tt.lessThan)
			}
		})
	}
}
