package healthcheck

type CoreWithCallback struct {
	itemOpts []itemCoreCallbackOpts
	core     *Core
}

func NewCoreWithCallbacks(opts ...itemCoreCallbackOpts) Health {
	return &CoreWithCallback{
		itemOpts: opts,
		core:     NewCore(),
	}
}

func (c *CoreWithCallback) Iterate(cb HealthIterator) error {
	return c.core.Iterate(cb)
}

func (c *CoreWithCallback) Get(name string) Item {
	c.core.lock.Lock()
	defer c.core.lock.Unlock()

	if _, ok := c.core.items[name]; !ok {
		c.core.items[name] = NewItemCoreWithCallbacks(name, c.itemOpts...)
	}

	return c.core.items[name]
}

func (c *CoreWithCallback) Stop(name string) {
	c.core.Stop(name)
}

func (c *CoreWithCallback) Status() map[string]bool {
	return c.core.Status()
}
