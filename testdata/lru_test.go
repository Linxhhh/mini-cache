package test

import (
	"mini-cache/lru"
	"testing"
)

// 定义一个新类型，实现 Value 接口的 Len 方法
type MyType string

func (t MyType) Len() int {
	return len(t) 
}

func TestLRU(t *testing.T) {
	cache := lru.New(int64(10))
	cache.Add("lxh", MyType("2005386"))

	if res, ok := cache.Get("lxh"); ok {
		t.Logf("value is %s\n", res)
		t.Log("Add() 和 Get() 方法有效！\n")
	} else {
		t.Log("Add() 或 Get() 方法存在错误！\n")
	}
}
