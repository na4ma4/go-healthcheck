package healthcheck

import (
	"net/http"
	"time"
)

type ReportStandard struct {
	Status   Status                `json:"status,omitempty"`
	Services []*ReportItemStandard `json:"services,omitempty"`
}

func (r *ReportStandard) GetStatus() string {
	return r.Status.String()
}

type ReportItemStandard struct {
	Name      string  `json:"name,omitempty"`
	Status    Status  `json:"status,omitempty"`
	StartTime string  `json:"start_time,omitempty"`
	Lifecycle []Event `json:"lifecycle,omitempty"`
}

func GenerateReport(healthCheck Health, r *http.Request) *ReportStandard {
	displayLifecycle := getQueryBool(r, "lifecycle")
	displayVerbose := getQueryBool(r, "verbose")

	out := &ReportStandard{
		Services: []*ReportItemStandard{},
	}

	_ = healthCheck.Iterate(func(name string, item Item) error {
		outItem := &ReportItemStandard{
			Name: name,
		}

		outItem.Status = item.Status()
		if out.Status.Less(outItem.Status) {
			out.Status = outItem.Status
		}

		if displayVerbose { // display verbose output (time)
			outItem.StartTime = item.StartTime().Format(time.RFC3339Nano)
		}

		if displayLifecycle || displayVerbose { // display lifecycle events
			outItem.Lifecycle = item.Lifecycle()
		}

		out.Services = append(out.Services, outItem)

		return nil
	})

	return out
}
