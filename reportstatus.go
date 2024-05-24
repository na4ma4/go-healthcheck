package healthcheck

type ReportStatus Status

const (
	ReportStatusGreen  ReportStatus = "green"
	ReportStatusYellow ReportStatus = "yellow"
	ReportStatusRed    ReportStatus = "red"
)

//nolint:gochecknoglobals // List of priority order for status.
var reportStatusList = []ReportStatus{
	ReportStatusRed,
	ReportStatusYellow,
	ReportStatusGreen,
}

func (s ReportStatus) String() string {
	return string(s)
}

func (s ReportStatus) Valid() bool {
	for _, v := range reportStatusList {
		if s == v {
			return true
		}
	}

	return false
}

// Less returns true if the status supplied is higher importance than the
// base Status.
func (s ReportStatus) Less(in ReportStatus) bool {
	return statusLess[ReportStatus](reportStatusList, s, in)
}

func ItemStatusToReportStatus(itemStatus Status) ReportStatus {
	switch itemStatus {
	case StatusRunning, StatusFinished:
		return ReportStatusGreen
	case StatusStarting:
		return ReportStatusYellow
	case StatusErrored, StatusUnknown:
		return ReportStatusRed
	}

	return ReportStatusRed
}
