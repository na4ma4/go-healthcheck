package healthcheck

import (
	"net/http"
	"time"
)

type ReportSimple struct {
	Status   ReportStatus        `json:"status,omitempty"`
	Services []*ReportItemSimple `json:"services,omitempty"`
}

func (r *ReportSimple) GetStatus() string {
	return r.Status.String()
}

type ReportItemSimple struct {
	Name      string       `json:"name,omitempty"`
	Status    ReportStatus `json:"status,omitempty"`
	StartTime string       `json:"start_time,omitempty"`
	Lifecycle []Event      `json:"lifecycle,omitempty"`
}

func GenerateReportSimple(healthCheck Health, r *http.Request) *ReportSimple {
	displayLifecycle := getQueryBool(r, "lifecycle")
	displayVerbose := getQueryBool(r, "verbose")

	out := &ReportSimple{
		Services: []*ReportItemSimple{},
	}

	out.Status = ReportStatusGreen

	_ = healthCheck.Iterate(func(name string, item Item) error {
		outItem := &ReportItemSimple{
			Name: name,
		}

		outItem.Status = ItemStatusToReportStatus(item.Status())
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
