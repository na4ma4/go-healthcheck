package healthcheck

import "time"

type Item interface {
	Name() string
	Duration() time.Duration
	Lifecycle() []Event
	StartTime() EventTime
	Error(err error) Item
	Start() Item
	Stop() Item
	Status() Status
	ToProto() *ItemProto
}
