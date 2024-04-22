package consistenthash

// 一致性哈希算法实现
// 针对同一 key 的查询，都会落到同一服务器节点上

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type hashFunc func([]byte) uint32

type HashRing struct {
	hash     hashFunc       // 哈希算法
	keys     []int          // 哈希环上的 key
	replicas int            // 虚拟节点，防止数据倾斜
	hashMap  map[int]string // 哈希值 -> 节点名
}

func NewHashRing(replicas int, fn hashFunc) *HashRing {
	r := &HashRing{
		hash:     fn,
		keys:     make([]int, 0),
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
	if fn == nil {
		r.hash = crc32.ChecksumIEEE   // 默认的哈希算法
	}
	return r
}

// 增加 key ，调整哈希环
func (hr *HashRing) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < hr.replicas; i++ {
			hash := int(hr.hash([]byte(strconv.Itoa(i) + key)))
			hr.keys = append(hr.keys, hash)
			hr.hashMap[hash] = key
		}
	}
	sort.Ints(hr.keys)
}

// 查询 key ，定位服务器节点
func (hr *HashRing) Get(key string) string {
	if len(hr.keys) == 0 {
		return ""
	}

	// 找到第一个大于等于 hash 值的节点
	hash := int(hr.hash([]byte(key)))
	idx := sort.Search(len(hr.keys), func(i int) bool {
		return hr.keys[i] >= hash
	})
	return hr.hashMap[hr.keys[idx % len(hr.keys)]]
}

