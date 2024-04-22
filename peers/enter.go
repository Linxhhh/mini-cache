package peers

// peers 模块，定义了获取分布式节点，以及从远端节点获取缓存的能力

// Picker 接口在 roles.Server 中实现
type Picker interface {
	Pick(key string) (peer Getter, ok bool)
}

// Getter 接口在 roles.client 中实现
type Getter interface {
	Get(group, key string) ([]byte, error)
}
