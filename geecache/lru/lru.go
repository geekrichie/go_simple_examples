package lru

import (
	"container/list"
)

type Cache struct {
	//允许使用的最大内存
	maxBytes int64
	//已经使用的内存
	nBytes   int64
	//
	ll *list.List
	cache map[string]*list.Element

	onEvicted func(key string, value Value)
}

type entry struct {
	key string
	value Value
}

type Value interface{
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache{
	return &Cache{
		maxBytes: maxBytes,
		ll : list.New(),
		cache: make(map[string]*list.Element),
		onEvicted: onEvicted,
	}
}

func (c *Cache)Get(key string) (value Value,ok bool){
	if elem, ok := c.cache[key];ok {
		c.ll.MoveToFront(elem)
		kv := elem.Value.(*entry)
		return kv.value, true
	}
    return
}

func (c *Cache) RemoveOldSet() {
	back := c.ll.Back()
	if back != nil{
		value := back.Value.(*entry)
		elemLen := len(value.key) + value.value.Len()
		c.ll.Remove(back)
		delete(c.cache, value.key)
		c.nBytes = c.nBytes-int64(elemLen)
		if c.onEvicted != nil {
			c.onEvicted(value.key,value.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		cur := ele.Value.(*entry)
		c.ll.MoveToFront(ele)
		c.nBytes += int64(value.Len()) - int64(cur.value.Len())
		cur.value = value
		//c.cache[key] = ele
	}else {
		c.ll.PushFront(&entry{key: key, value:value})
		c.cache[key] = c.ll.Front()
		c.nBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldSet()
	}
}

func (c *Cache) Len() int{
	return c.ll.Len()
}