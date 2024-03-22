package health

type Status string

const (
	StatusStarting Status = "starting"
	StatusRunning  Status = "running"
	StatusFinished Status = "finished"
	StatusErrored  Status = "errored"
	StatusUnknown  Status = "unknown"
)

func StatusIsRunning(s Status) bool {
	return s == StatusRunning
}
