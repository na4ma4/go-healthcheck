package health

import "time"

type Item interface {
	Name() string
	Duration() time.Duration
	Lifecycle() []Event
	StartTime() time.Time
	Error(err error) Item
	Start() Item
	Stop() Item
	Status() Status
}
