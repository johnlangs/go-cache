package cache

import (
	"sync"
	"time"
)

type tableEntry struct {
	item     interface{}
	lifetime int
}

type cache struct {
	sync.RWMutex
	table          map[string]*tableEntry
	stdTTL         int
	checkInterval  int
	deleteOnExpire bool
	maxKeys        int
	currentKeys    int
}

func CreateCache(stdTTL int, checkInterval int, deleteOnExpire bool, maxKeys int) *cache {
	c := &cache{
		table:          make(map[string]*tableEntry),
		stdTTL:         stdTTL,
		checkInterval:  checkInterval,
		deleteOnExpire: deleteOnExpire,
		maxKeys:        maxKeys,
		currentKeys:    0,
	}

	if deleteOnExpire {
		go c.lifetimeWatcher()
	}

	return c
}

func (c *cache) lifetimeWatcher() {
	for {
		time.Sleep(time.Duration(c.checkInterval) * time.Second)

		c.Lock()
		for key, entry := range c.table {
			entry.lifetime += c.checkInterval
			if entry.lifetime >= c.stdTTL {
				delete(c.table, key)
				c.currentKeys = len(c.table)
			}
		}
		c.Unlock()
	}
}

func (c *cache) Set(key string, item interface{}) bool {
	entry := &tableEntry{
		item: item,
	}

	c.Lock()

	if _, ok := c.table[key]; ok {
		c.table[key] = entry
		c.Unlock()
		return true
	}

	if c.currentKeys+1 > c.maxKeys && c.maxKeys > 0 {
		c.Unlock()
		return false
	}

	c.table[key] = entry
	c.currentKeys++

	c.Unlock()
	return true
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
	c.currentKeys = len(c.table)
	c.Unlock()
}
