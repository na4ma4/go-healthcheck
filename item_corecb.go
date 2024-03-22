package health

import (
	"time"
)

type ItemCallbackFunc func(status Status, item Item, err error) (continueExecutuion bool)

type itemCoreCallbackOpts func(itemCore *ItemCoreWithCallback)

type ItemCoreWithCallback struct {
	cb   map[Status]ItemCallbackFunc
	item *ItemCore
}

func NewItemCoreWithCallbacks(name string, opts ...itemCoreCallbackOpts) *ItemCoreWithCallback {
	item := &ItemCoreWithCallback{
		cb:   map[Status]ItemCallbackFunc{},
		item: NewItemCore(name),
	}

	for _, opt := range opts {
		opt(item)
	}

	if !item.runCallback(StatusStarting, nil) {
		item.item.times = map[Status]time.Time{}
		item.item.lifecycle = []Event{}
	}

	return item
}

func (i *ItemCoreWithCallback) Name() string {
	return i.item.Name()
}

func (i *ItemCoreWithCallback) runCallback(status Status, err error) bool {
	if f, ok := i.cb[status]; ok {
		return f(status, i.item, err)
	}

	return true
}

func (i *ItemCoreWithCallback) Duration() time.Duration {
	return i.item.Duration()
}

func (i *ItemCoreWithCallback) StartTime() time.Time {
	return i.item.StartTime()
}

func (i *ItemCoreWithCallback) Error(err error) Item {
	if !i.runCallback(StatusErrored, err) {
		return i
	}
	return i.item.Error(err)
}

func (i *ItemCoreWithCallback) Start() Item {
	if !i.runCallback(StatusRunning, nil) {
		return i
	}
	return i.item.Start()
}

func (i *ItemCoreWithCallback) Stop() Item {
	if !i.runCallback(StatusFinished, nil) {
		return i
	}
	return i.item.Stop()
}

func (i *ItemCoreWithCallback) Status() Status {
	return i.item.Status()
}

func (i *ItemCoreWithCallback) Lifecycle() []Event {
	return i.item.Lifecycle()
}

//nolint:revive // unexported so option functions aren't created outside the module.
func AddOnStatusCallback(status Status, cb ItemCallbackFunc) itemCoreCallbackOpts {
	return func(itemCore *ItemCoreWithCallback) {
		itemCore.cb[status] = cb
	}
}
