package healthcheck

// type Status string

const (
	StatusStarting = Status_STARTING
	StatusRunning  = Status_RUNNING
	StatusFinished = Status_FINISHED
	StatusErrored  = Status_ERRORED
	StatusUnknown  = Status_UNKNOWN
)

//nolint:gochecknoglobals // List of priority order for status.
var statusPriorityList = []Status{
	Status_ERRORED,
	Status_UNKNOWN,
	Status_STARTING,
	Status_RUNNING,
	Status_FINISHED,
}

func StatusIsHealthy(s Status) bool {
	return s == StatusRunning || s == StatusFinished
}

func (s Status) Valid() bool {
	for _, v := range statusPriorityList {
		if s == v {
			return true
		}
	}

	return false
}

// Less returns true if the status supplied is higher importance than the
// base Status.
func (s Status) Less(in Status) bool {
	return statusLess[Status](statusPriorityList, s, in)
}
