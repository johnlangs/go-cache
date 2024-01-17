A naive implementation of a cache in Go.

Useage:

  c := NewCache(stdTTL int, checkInterval int, deleteOnExpire bool, maxKeys int)
  
  cache.Set(key string, item interface{})
  
  cache.Get(key string)
  
  cache.Delete(key string)
