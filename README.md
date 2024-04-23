### 项目简介

mini-cache 项目的整体结构大致如下：

![image.png](https://flowus.cn/preview/6a62cc3e-ac33-4254-b81b-502f43fd20b3)



**技术内容：etcd + gRPC + LRU + singleflight + consistenthash**

- 使用 SingleFlight 缓存模式，防止高并发场景下的缓存击穿问题；

- 使用一致性哈希算法，使得同一 Key 的查询都会落在同一节点上；

- 使用 Etcd 实现服务注册和服务发现，服务节点间通过 gRPC 通信。



### 项目使用

下载安装：

```Go
go get -u github.com/Linxhhh/mini-cache
```




代码示例：

```Go
package main

import (
	"fmt"
	"log"
	"sync"

	mini_cache "github.com/Linxhhh/mini-cache"
)

// 模拟数据库
var DB = map[string]string{
	"1": "lxh",
	"2": "lxt",
	"3": "lxj",
}

func main() {
	// 创建缓存分组
	stuGroup := mini_cache.NewGroup("student", mini_cache.GetterFunc(func(key string) ([]byte, error) {
		log.Printf("Search DB for key [%s]\n", key)
		if val, ok := DB[key]; ok {
			return []byte(val), nil
		}
		return nil, fmt.Errorf("can't find key [%s] in DB, err : %s", key)
	}), 2<<10)
	
	// 创建服务节点
	addr := "127.0.0.1:3333"
	server, err := mini_cache.NewServer(addr)
	if err != nil {
		log.Fatalf("create server failed, err : %s", err)
	}
	server.SetPeers(addr)

	// 缓存分组绑定服务节点
	stuGroup.RegisterPeer(server)
	
	// 服务节点上线
	go func() {
		err = server.Start()
		if err != nil {
			log.Fatal("server start failed, err : %s", err)
		}
	}()

	// 发出几个Get请求
	var wg sync.WaitGroup
	wg.Add(4)
	go GetStudent(stuGroup, &wg)
	go GetStudent(stuGroup, &wg)
	go GetStudent(stuGroup, &wg)
	go GetStudent(stuGroup, &wg)
	wg.Wait()
}

func GetStudent(group *mini_cache.Group, wg *sync.WaitGroup) {
	log.Printf("get Student ...")
	view, err := group.Get("1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(view.String())
	wg.Done()
}

/*
2024/04/23 20:09:42 get Student ...
2024/04/23 20:09:42 [server 127.0.0.1:3333] pick itself, key is 1
2024/04/23 20:09:42 Search DB for key [1]
lxh
2024/04/23 20:09:42 get Student ...
2024/04/23 20:09:42 cache hit!
2024/04/23 20:09:42 get Student ...
lxh
2024/04/23 20:09:42 get Student ...
2024/04/23 20:09:42 cache hit!
lxh
2024/04/23 20:09:42 cache hit!
lxh
*/
```




