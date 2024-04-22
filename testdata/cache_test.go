package test

import (
	"fmt"
	"log"
	mini_cache "mini-cache"
	"testing"
)

// 模拟数据库
var db = map[string]string{
	"math":    "100",
	"English": "99",
	"Chinese": "98",
}

func TestGroup(t *testing.T) {

	// 回调函数，从 db 中加载数据
	getter := mini_cache.GetterFunc(func(key string) ([]byte, error) {
		log.Printf("[slowdb select] for [key:%s]", key)
		for k, v := range db {
			if k == key {
				return []byte(v), nil
			}
		}
		return nil, fmt.Errorf("[key:%s] is not found", key)
	})

	scoreGroup := mini_cache.NewGroup("scores", getter, int64(1<<20))

	// 查询缓存(第一次未命中，第二次命中)
	byteview, err := scoreGroup.Get("math")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(byteview)
	byteview, err = scoreGroup.Get("math")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(byteview)

	// 查询不存在的数据
	byteview, err = scoreGroup.Get("science")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(byteview)
}
