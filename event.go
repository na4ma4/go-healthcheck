package health

import (
	"fmt"
	"time"
)

type Event struct {
	Timestamp time.Time
	Status    Status
}

func (h Event) String() string {
	return fmt.Sprintf("%s[%s]", h.Status, h.Timestamp.Format(time.RFC3339))
}
