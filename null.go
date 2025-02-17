package healthcheck

import "time"

const (
	nullItemName = "null"
)

//nolint:gochecknoglobals // null item values to reduce allocations.
var (
	nullLifecycle = []Event{}
	nullDuration  = time.Duration(0)
	nullStartTime = NewEventTime(time.Time{})
	nullItem      = &NullItem{}
)

type Null struct {
}

func NewNull() *Null {
	return &Null{}
}

func (n *Null) Get(_ string) Item               { return nullItem }
func (n *Null) Stop(_ string)                   {}
func (n *Null) Status() map[string]bool         { return map[string]bool{} }
func (n *Null) Iterate(cb HealthIterator) error { return cb(nullItemName, nullItem) }
func (n *Null) ToProto() *CoreProto             { return &CoreProto{} }

type NullItem struct{}

func (n *NullItem) Name() string            { return nullItemName }
func (n *NullItem) Lifecycle() []Event      { return nullLifecycle }
func (n *NullItem) Duration() time.Duration { return nullDuration }
func (n *NullItem) StartTime() EventTime    { return nullStartTime }
func (n *NullItem) Error(_ error) Item      { return n }
func (n *NullItem) Start() Item             { return n }
func (n *NullItem) Stop() Item              { return n }
func (n *NullItem) Status() Status          { return StatusUnknown }
func (n *NullItem) ToProto() *ItemProto     { return &ItemProto{} }
