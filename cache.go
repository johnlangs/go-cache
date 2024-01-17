package cache

import "sync"

type tableEntry struct {
	item interface{}
	lifetime int
}

type cache struct {
	sync.RWMutex
	stdTTL int
	table map[string]*tableEntry
	deleteOnExpire bool
	maxKeys int
}

func CreateCache(stdTTL int, maxKeys int, deleteOnExpire bool) *cache {
	return &cache{
		table: make(map[string]*tableEntry),
		stdTTL: stdTTL,
		maxKeys: maxKeys,
		deleteOnExpire: deleteOnExpire,
	}
}

func (c *cache) Set(key string, item interface{}) error {
	entry := &tableEntry{
		item: item,
	}

	c.Lock()
	c.table[key] = entry
	c.Unlock()

	return nil
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.Lock()

	entry, ok := c.table[key]
	if !ok {
		c.Unlock()
		return nil, false
	}

	c.Unlock()
	return entry.item, true
}

func (c *cache) Delete(key string) {
	c.Lock()
	delete(c.table, key)
	c.Unlock()
}
