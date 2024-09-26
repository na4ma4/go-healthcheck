package healthcheck

import (
	"net/http"
	"time"
)

type ReportSimple struct {
	status   ReportStatus        `json:"-"`
	Status   ReportStatusText    `json:"status,omitempty"`
	Services []*ReportItemSimple `json:"services,omitempty"`
}

func (r *ReportSimple) GetStatus() string {
	return r.Status.String()
}

type ReportItemSimple struct {
	Name      string           `json:"name,omitempty"`
	status    ReportStatus     `json:"-"`
	Status    ReportStatusText `json:"status,omitempty"`
	StartTime string           `json:"start_time,omitempty"`
	Lifecycle []Event          `json:"lifecycle,omitempty"`
}

func GenerateReportSimple(healthCheck Health, r *http.Request) *ReportSimple {
	displayLifecycle := getQueryBool(r, "lifecycle")
	displayVerbose := getQueryBool(r, "verbose")

	out := &ReportSimple{
		Services: []*ReportItemSimple{},
	}

	out.status = ReportStatusGreen

	_ = healthCheck.Iterate(func(name string, item Item) error {
		outItem := &ReportItemSimple{
			Name: name,
		}

		outItem.status = ItemStatusToReportStatus(item.Status())
		if out.status.Less(outItem.status) {
			out.status = outItem.status
		}

		outItem.Status = ReportStatusTextFromStatus(outItem.status)

		if displayVerbose { // display verbose output (time)
			outItem.StartTime = item.StartTime().Format(time.RFC3339Nano)
		}

		if displayLifecycle || displayVerbose { // display lifecycle events
			outItem.Lifecycle = item.Lifecycle()
		}

		out.Services = append(out.Services, outItem)

		return nil
	})

	out.Status = ReportStatusTextFromStatus(out.status)

	return out
}
