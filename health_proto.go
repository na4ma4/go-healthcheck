package healthcheck

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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
			Status:    event.Status,
		})
	}

	i.lock.Lock()
	defer i.lock.Unlock()
	for status, ts := range i.times {
		o.Times[int32(status)] = timestamppb.New(ts.Time)
	}

	return o
}
