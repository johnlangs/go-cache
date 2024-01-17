package cache

type cache struct {
	table map[string]interface{}
	
}

func CreateCache() cache {
	var c cache
	c.table = make(map[string]interface{})
	return c
}

func (c cache) Set(key string, item interface{}) error {
	c.table[key] = item
	return nil
}

func (c cache) Get(key string) (interface{}, bool) {
	var item interface{}
	var ok bool
	item, ok = c.table[key]
	return item, ok
}
