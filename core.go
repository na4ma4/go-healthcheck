package healthcheck

import (
	"sync"
)

type Core struct {
	lock  sync.RWMutex
	items map[string]Item
}

func NewCore() *Core {
	return &Core{
		items: map[string]Item{},
	}
}

func (c *Core) Get(name string) Item {
	c.lock.Lock()
	defer c.lock.Unlock()

	if _, ok := c.items[name]; !ok {
		c.items[name] = NewItemCore(name)
	}

	return c.items[name]
}

func (c *Core) Stop(name string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if _, ok := c.items[name]; ok {
		c.items[name].Stop()
	}
}

func (c *Core) Status() map[string]bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	out := map[string]bool{}

	for k, v := range c.items {
		out[k] = StatusIsRunning(v.Status())
	}

	return out
}
