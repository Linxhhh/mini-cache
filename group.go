package mini_cache

import (
	"fmt"
	"log"
	"github.com/Linxhhh/mini-cache/singleflight"
	"github.com/Linxhhh/mini-cache/peers"
	"sync"
)

// Group 可视作缓存的命名空间，每个 Group 具有唯一的 name
// 比如，缓存学生的成绩命名为 scores，缓存学生信息的命名为 info，缓存学生课程的命名为 courses

type Group struct {
	name   string
	getter Getter
	cache  Cache

	peer   peers.Picker
	flight *singleflight.Flight
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

// 新建缓存分组
func NewGroup(name string, getter Getter, cacheByte int64) *Group {
	if getter == nil {
		log.Panic("not getter!")
	}

	mu.Lock()
	defer mu.Unlock()

	group := &Group{
		name:   name,
		getter: getter,
		cache:  *NewCache(cacheByte),
		flight: &singleflight.Flight{},
	}
	groups[name] = group

	return group
}

// 为分组注册服务器节点
func (g *Group) RegisterPeer(peer peers.Picker) {
	if g.peer != nil {
		panic("group had been registered server")
	}
	g.peer = peer
}

// 获取分组实例
func GetGroup(name string) (*Group, bool) {
	mu.RLock()
	g, ok := groups[name]
	mu.RUnlock()
	return g, ok
}

// 从当前的缓存分组中获取数据
func (g *Group) Get(key string) (value Byteview, err error) {
	if key == "" {
		return Byteview{}, fmt.Errorf("key is empty！")
	}

	if b, ok := g.cache.Get(key); ok {
		log.Println("cache hit!")
		return b, nil
	}

	return g.load(key)
}

// 当缓存未命中时，从其它数据源加载数据
func (g *Group) load(key string) (value Byteview, err error) {
	val, err := g.flight.Do(key, func() (interface{}, error) {
		if g.peer != nil {
			if peer, ok := g.peer.Pick(key); ok {
				if val, err := g.getFromPeer(peer, key); err == nil {
					return val, err
				}
				log.Println("[mini-Cache] Failed to get from peer", err)
			}
		}
		return g.getlocally(key)
	})
	
	if err == nil {
		return val.(Byteview), err
	}
	return
}

// 从其他远程节点中，获取缓存数据
func (g *Group) getFromPeer(peer peers.Getter, key string) (Byteview, error) {
	value, err := peer.Get(g.name, key)
	if err != nil {
		return Byteview{}, fmt.Errorf("get from peer failed, err : %s", err)
	}
	return Byteview{b: value}, nil
}

// 使用 getter 回调函数，将（数据库）数据载入缓存中
func (g *Group) getlocally(key string) (Byteview, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return Byteview{}, fmt.Errorf("get locally failed, err : %s", err)
	}

	value := Byteview{b: bytes}
	g.cache.Add(key, value)
	return value, nil
}
