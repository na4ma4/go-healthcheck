package healthcheck

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type healthReportStatus string

func (h healthReportStatus) String() string {
	return string(h)
}

const (
	statusGreen  healthReportStatus = "green"
	statusYellow healthReportStatus = "yellow"
	statusRed    healthReportStatus = "red"
)

type healthReport struct {
	Status   healthReportStatus  `json:"status,omitempty"`
	Services []*healthReportItem `json:"services,omitempty"`
}

type healthReportItem struct {
	Name      string             `json:"name,omitempty"`
	Status    healthReportStatus `json:"status,omitempty"`
	StartTime string             `json:"start_time,omitempty"`
	Lifecycle []Event            `json:"lifecycle,omitempty"`
}

func Handler(healthCheck Health) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		report := healthToReport(healthCheck, r)
		w.Header().Add("X-Health-Status", report.Status.String())
		_ = json.NewEncoder(w).Encode(report)
	}
}

func healthItemStatusToReportStatus(itemStatus Status) healthReportStatus {
	switch itemStatus {
	case StatusRunning:
		return statusGreen
	case StatusStarting, StatusFinished:
		return statusYellow
	case StatusErrored, StatusUnknown:
		return statusRed
	}

	return statusRed
}

func getQueryBool(r *http.Request, name string) bool {
	v := r.URL.Query().Get(name)
	switch strings.ToLower(v) {
	case "yes", "true", "1", "on":
		return true
	}

	return false
}

func healthToReport(healthCheck Health, r *http.Request) *healthReport {
	out := &healthReport{
		Status:   statusGreen,
		Services: make([]*healthReportItem, 0),
	}

	displayLifecycle := getQueryBool(r, "lifecycle")
	displayVerbose := getQueryBool(r, "verbose")

	_ = healthCheck.Iterate(func(name string, item Item) error {
		outItem := &healthReportItem{
			Name:   name,
			Status: healthItemStatusToReportStatus(item.Status()),
		}

		if displayVerbose {
			outItem.StartTime = item.StartTime().Format(time.RFC3339Nano)
		}

		if displayLifecycle || displayVerbose {
			outItem.Lifecycle = item.Lifecycle()
		}

		if outItem.Status != statusGreen && (out.Status == statusYellow || out.Status == statusGreen) {
			out.Status = outItem.Status
		}

		out.Services = append(out.Services, outItem)

		return nil
	})

	return out
}
