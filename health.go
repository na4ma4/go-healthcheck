package healthcheck

type Health interface {
	Get(name string) Item
	Stop(name string)
	Status() map[string]bool
}

// func LifecycleTimeToDuration(lifecycle map[time.Time]Status) map[time.Duration]Status {
// 	out := map[time.Duration]Status{}

// 	if len(lifecycle) == 0 {
// 		return out
// 	}

// 	keys := []time.Time{}
// 	for ts := range lifecycle {
// 		keys = append(keys, ts)
// 	}

// 	sort.Slice(keys, func(i, j int) bool {
// 		return !keys[i].After(keys[j])
// 	})

// 	startTime := keys[0]

// 	for ts, status := range lifecycle {
// 		out[ts.Sub(startTime)] = status
// 	}

// 	return out
// }
