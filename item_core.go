package healthcheck

import (
	"slices"
	"sync"
	"time"
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
			{NewEventTime(ts), StatusStarting},
		},
	}
}

func (h *ItemCore) setStatus(s Status) {
	ts := time.Now()
	h.status = s
	h.times[s] = NewEventTime(ts)
	h.lifecycle = append(h.lifecycle, Event{NewEventTime(ts), s})
}

func (h *ItemCore) Name() string {
	return h.name
}

func (h *ItemCore) SetStatus(s Status) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.setStatus(s)
}

func (h *ItemCore) Duration() time.Duration {
	h.lock.RLock()
	defer h.lock.RUnlock()

	for _, status := range []Status{StatusFinished, StatusErrored} {
		if v, ok := h.times[status]; ok {
			return v.Sub(h.times[StatusStarting].Time)
		}
	}

	return time.Since(h.times[StatusStarting].Time)
}

func (h *ItemCore) Lifecycle() []Event {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return slices.Clone(h.lifecycle)
}

func (h *ItemCore) StartTime() EventTime {
	h.lock.Lock()
	defer h.lock.Unlock()

	return h.times[StatusStarting]
}

func (h *ItemCore) Status() Status {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return h.status
}

func (h *ItemCore) Start() Item {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.setStatus(StatusRunning)
	return h
}

func (h *ItemCore) Error(err error) Item {
	h.lock.Lock()
	defer h.lock.Unlock()

	if h.status == StatusRunning || h.status == StatusStarting {
		if err != nil {
			h.setStatus(StatusErrored)
		} else {
			h.setStatus(StatusFinished)
		}
	}

	return h
}

func (h *ItemCore) Stop() Item {
	h.lock.Lock()
	defer h.lock.Unlock()

	if h.status == StatusRunning || h.status == StatusStarting {
		h.setStatus(StatusFinished)
	}

	return h
}
