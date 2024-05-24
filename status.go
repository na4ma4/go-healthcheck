package healthcheck

type Status string

const (
	StatusStarting Status = "starting"
	StatusRunning  Status = "running"
	StatusFinished Status = "finished"
	StatusErrored  Status = "errored"
	StatusUnknown  Status = "unknown"
)

//nolint:gochecknoglobals // List of priority order for status.
var statusList = []Status{
	StatusErrored,
	StatusUnknown,
	StatusStarting,
	StatusRunning,
	StatusFinished,
}

func StatusIsHealthy(s Status) bool {
	return s == StatusRunning || s == StatusFinished
}

func (s Status) String() string {
	return string(s)
}

func (s Status) Valid() bool {
	for _, v := range statusList {
		if s == v {
			return true
		}
	}

	return false
}

// Less returns true if the status supplied is higher importance than the
// base Status.
func (s Status) Less(in Status) bool {
	return statusLess[Status](statusList, s, in)
}

func statusLess[T ~string](src []T, l, r T) bool {
	if l == "" {
		return true
	}
	if l == r {
		return false
	}
	for _, v := range src {
		if v == l {
			return false
		} else if v == r {
			return true
		}
	}

	return true
}
