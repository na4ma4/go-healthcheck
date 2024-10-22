package healthcheck

import (
	"slices"
	"sync"
	"time"

	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type ItemCore struct {
	lock      sync.RWMutex
	name      string
	status    Status
	times     map[Status]EventTime
	lifecycle []Event
}

func NewItemCore(name string) *ItemCore {
	ts := time.Now()
	return &ItemCore{
		name:   name,
		status: StatusStarting,
		times: map[Status]EventTime{ //nolint:exhaustive // initial state.
			StatusStarting: NewEventTime(ts),
		},
		lifecycle: []Event{
			{NewEventTime(ts), NewEventStatus(StatusStarting)},
		},
	}
}

func (i *ItemCore) setStatus(s Status) {
	ts := time.Now()
	i.status = s
	i.times[s] = NewEventTime(ts)
	i.lifecycle = append(i.lifecycle, Event{NewEventTime(ts), NewEventStatus(s)})
}

func (i *ItemCore) Name() string {
	return i.name
}

func (i *ItemCore) SetStatus(s Status) {
	i.lock.Lock()
	defer i.lock.Unlock()

	i.setStatus(s)
}

func (i *ItemCore) Duration() time.Duration {
	i.lock.RLock()
	defer i.lock.RUnlock()

	for _, status := range []Status{StatusFinished, StatusErrored} {
		if v, ok := i.times[status]; ok {
			return v.Sub(i.times[StatusStarting].Time)
		}
	}

	return time.Since(i.times[StatusStarting].Time)
}

func (i *ItemCore) Lifecycle() []Event {
	i.lock.RLock()
	defer i.lock.RUnlock()

	return slices.Clone(i.lifecycle)
}

func (i *ItemCore) StartTime() EventTime {
	i.lock.Lock()
	defer i.lock.Unlock()

	return i.times[StatusStarting]
}

func (i *ItemCore) Status() Status {
	i.lock.RLock()
	defer i.lock.RUnlock()

	return i.status
}

func (i *ItemCore) Start() Item {
	i.lock.Lock()
	defer i.lock.Unlock()

	i.setStatus(StatusRunning)
	return i
}

func (i *ItemCore) Error(err error) Item {
	i.lock.Lock()
	defer i.lock.Unlock()

	if i.status == StatusRunning || i.status == StatusStarting {
		if err != nil {
			i.setStatus(StatusErrored)
		} else {
			i.setStatus(StatusFinished)
		}
	}

	return i
}

func (i *ItemCore) Stop() Item {
	i.lock.Lock()
	defer i.lock.Unlock()

	if i.status == StatusRunning || i.status == StatusStarting {
		i.setStatus(StatusFinished)
	}

	return i
}

func (i *ItemCore) ToProto() *ItemProto {
	o := &ItemProto{
		Name:      i.Name(),
		Duration:  durationpb.New(i.Duration()),
		StartTime: timestamppb.New(i.StartTime().Time),
		Status:    i.Status(),
		Lifecycle: make([]*EventProto, 0),
		Times:     make(map[int32]*timestamppb.Timestamp),
	}

	for _, event := range i.Lifecycle() {
		o.Lifecycle = append(o.Lifecycle, &EventProto{
			Timestamp: timestamppb.New(event.Timestamp.Time),
			Status:    event.Status.Status,
		})
	}

	i.lock.Lock()
	defer i.lock.Unlock()
	for status, ts := range i.times {
		o.Times[int32(status)] = timestamppb.New(ts.Time)
	}

	return o
}
