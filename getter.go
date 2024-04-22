package mini_cache

// 回调函数，用于加载数据（一般是从数据库加载）

type Getter interface{
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (g GetterFunc) Get(key string) ([]byte, error) {
	return g(key)
}