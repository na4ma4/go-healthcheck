package healthcheck

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Timestamp EventTime   `json:"ts"`
	Status    EventStatus `json:"status"`
}

func (h Event) String() string {
	return fmt.Sprintf("%s[%s]", h.Status, h.Timestamp.Format(time.RFC3339))
}

func NewEventStatus(st Status) EventStatus {
	return EventStatus{st}
}

type EventStatus struct {
	Status
}

func (s EventStatus) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s EventStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.String() + `"`), nil
}

func (s *EventStatus) UnmarshalJSON(data []byte) error {
	var in string
	if err := json.Unmarshal(data, &in); err != nil {
		return err
	}

	if v, ok := Status_value[in]; ok {
		*s = EventStatus{Status(v)}
		return nil
	}

	*s = NewEventStatus(Status_UNKNOWN)
	return nil
}

func NewEventTime(ts time.Time) EventTime {
	return EventTime{ts}
}

type EventTime struct {
	time.Time
}

func (j EventTime) format() string {
	return j.Time.Format(time.RFC3339Nano)
}

func (j EventTime) MarshalText() ([]byte, error) {
	return []byte(j.format()), nil
}

func (j EventTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + j.format() + `"`), nil
}
