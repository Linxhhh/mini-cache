package lru

import "container/list"

type Cache struct {
	ll       *list.List
	cache    map[string]*list.Element
	curBytes int64
	maxBytes int64
}

// list.element 以 entry 的形式存放
type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64) *Cache {
	return &Cache{
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
		curBytes: 0,
		maxBytes: maxBytes,
	}
}

func (lru *Cache) Get(key string) (Value, bool) {
	if element, ok := lru.cache[key]; ok {
		lru.ll.MoveToBack(element)
		kv := element.Value.(*entry)
		return kv.value, true
	} 
	return nil, false 
}

func (lru *Cache) Add(key string, value Value) {
	if element, ok := lru.cache[key]; ok {
		lru.ll.MoveToBack(element)
		kv := element.Value.(*entry)
		lru.curBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		element := lru.ll.PushBack(&entry{key, value})
		lru.cache[key] = element
		lru.curBytes += int64(len(key)) + int64(value.Len())
	}

	// 如果字节数已满，则淘汰
	for lru.curBytes > lru.maxBytes {
		lru.RemoveOldest()
	}
}

func (lru *Cache) RemoveOldest() {
	if lru.ll.Len() > 0 {
		kv := lru.ll.Back().Value.(*entry)
		delete(lru.cache, kv.key)
		lru.ll.Remove(lru.ll.Back())
		lru.curBytes -= int64(len(kv.key)) + int64(kv.value.Len())
	}
}

func (lru *Cache) Len() int {
	return lru.ll.Len()
}