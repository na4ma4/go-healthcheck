package healthcheck

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Handler(healthCheck Health) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		var report Report
		{
			if getQueryBool(r, "simple") {
				report = GenerateReportSimple(healthCheck, r)
			} else {
				report = GenerateReport(healthCheck, r)
			}
		}
		w.Header().Add("X-Health-Status", report.GetStatus())
		_ = json.NewEncoder(w).Encode(report)
	}
}

func getQueryBool(r *http.Request, name string) bool {
	v := r.URL.Query().Get(name)
	switch strings.ToLower(v) {
	case "yes", "true", "1", "on":
		return true
	}

	return false
}

// func healthToReport[S reportStatus, P iReportStatus[S]](healthCheck Health, r *http.Request) T {
// 	displayLifecycle := getQueryBool(r, "lifecycle")
// 	displayVerbose := getQueryBool(r, "verbose")
// 	displaySimple := getQueryBool(r, "simple")

// 	out := &Report[S, P]{
// 		Services: []*ReportItem[S, P]{},
// 	}

// 	if displaySimple {
// 		out.Status = S(ReportStatusGreen)
// 	} else {
// 		out.Status = S(StatusRunning)
// 	}

// 	_ = healthCheck.Iterate(func(name string, item Item) error {
// 		outItem := &ReportItem[S, P]{
// 			Name: name,
// 		}

// 		if displaySimple { // display simple (green/yellow/red) or detailed status
// 			outItem.Status = S(ItemStatusToReportStatus(item.Status()))
// 			if out.Status.Less(outItem.Status) {
// 				out.Status = outItem.Status
// 			}
// 		} else { // detailed status
// 			outItem.Status = ReportStatus(item.Status())
// 			if out.Status.Less(outItem.Status) {
// 				out.Status = outItem.Status
// 			}
// 		}

// 		if displayVerbose { // display verbose output (time)
// 			outItem.StartTime = item.StartTime().Format(time.RFC3339Nano)
// 		}

// 		if displayLifecycle || displayVerbose { // display lifecycle events
// 			outItem.Lifecycle = item.Lifecycle()
// 		}

// 		out.Services = append(out.Services, outItem)

// 		return nil
// 	})

// 	return out
// }

// func healthToReportSimple(healthCheck Health, r *http.Request) *ReportSimple {

// }

// func healthToReport(healthCheck Health, r *http.Request) *Report {
// 	displayLifecycle := getQueryBool(r, "lifecycle")
// 	displayVerbose := getQueryBool(r, "verbose")

// 	out := &Report{
// 		Services: make([]*ReportItem, 0),
// 	}

// 	if displaySimple {
// 		out.Status = ReportStatusGreen
// 	} else {
// 		out.Status = ReportStatus(StatusRunning)
// 	}

// 	_ = healthCheck.Iterate(func(name string, item Item) error {
// 		outItem := &ReportItem{
// 			Name: name,
// 		}

// 		if displaySimple { // display simple (green/yellow/red) or detailed status
// 			outItem.Status = ItemStatusToReportStatus(item.Status())
// 			if out.Status.Less(outItem.Status) {
// 				out.Status = outItem.Status
// 			}
// 		} else { // detailed status
// 			outItem.Status = ReportStatus(item.Status())
// 			if out.Status.Less(outItem.Status) {
// 				out.Status = outItem.Status
// 			}
// 		}

// 		if displayVerbose { // display verbose output (time)
// 			outItem.StartTime = item.StartTime().Format(time.RFC3339Nano)
// 		}

// 		if displayLifecycle || displayVerbose { // display lifecycle events
// 			outItem.Lifecycle = item.Lifecycle()
// 		}

// 		out.Services = append(out.Services, outItem)

// 		return nil
// 	})

// 	return out
// }
