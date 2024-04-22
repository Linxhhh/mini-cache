package singleflight

// 缓存模式 singleflight
// 确保在同一时刻，对于同一个 key 的多个查询请求，如果缓存未命中，只有一个请求会真正地去查询数据库

import "sync"

type visitDB struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Flight struct {
	mu     sync.Mutex
	flight map[string]*visitDB
}

func (f *Flight) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	// 初始化、条件检查
	f.mu.Lock()
	if f.flight == nil {
		f.flight = make(map[string]*visitDB)
	}
	if v, ok := f.flight[key]; ok {
		f.mu.Unlock()
		v.wg.Wait()     // wait 确保 val 和 err 已取到
		return v.val, v.err
	}
	v := new(visitDB)
	v.wg.Add(1)
	f.flight[key] = v
	f.mu.Unlock()

	v.val, v.err = fn()
	v.wg.Done()

	f.mu.Lock()
	delete(f.flight, key)
	f.mu.Unlock()

	return v.val, v.err
}
