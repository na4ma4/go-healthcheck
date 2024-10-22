package healthcheck

type HealthIterator func(name string, item Item) error

type Health interface {
	Iterate(cb HealthIterator) error
	ToProto() *CoreProto
	Get(name string) Item
	Stop(name string)
	Status() map[string]bool
}
