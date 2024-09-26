package healthcheck

import (
	"net/http"
	"time"
)

type ReportStandard struct {
	status   *Status               `json:"-"`
	Status   StatusText            `json:"status,omitempty"`
	Services []*ReportItemStandard `json:"services,omitempty"`
}

func (r *ReportStandard) GetStatus() string {
	return r.Status.String()
}

type ReportItemStandard struct {
	Name      string     `json:"name,omitempty"`
	status    *Status    `json:"-"`
	Status    StatusText `json:"status,omitempty"`
	StartTime string     `json:"start_time,omitempty"`
	Lifecycle []Event    `json:"lifecycle,omitempty"`
}

func GenerateReport(healthCheck Health, r *http.Request) *ReportStandard {
	displayLifecycle := getQueryBool(r, "lifecycle")
	displayVerbose := getQueryBool(r, "verbose")

	out := &ReportStandard{
		Status:   StatusText(Status_UNKNOWN.String()),
		Services: []*ReportItemStandard{},
	}

	_ = healthCheck.Iterate(func(name string, item Item) error {
		outItem := &ReportItemStandard{
			Name: name,
		}

		itemStatus := item.Status()
		outItem.status = &itemStatus
		if out.status == nil || out.status.Less(itemStatus) {
			out.status = outItem.status
		}

		outItem.Status = StatusTextFromStatus(*outItem.status)

		if displayVerbose { // display verbose output (time)
			outItem.StartTime = item.StartTime().Format(time.RFC3339Nano)
		}

		if displayLifecycle || displayVerbose { // display lifecycle events
			outItem.Lifecycle = item.Lifecycle()
		}

		out.Services = append(out.Services, outItem)

		return nil
	})

	if out.status != nil {
		out.Status = StatusTextFromStatus(*out.status)
	}

	return out
}
