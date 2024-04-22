package mini_cache

// 定义 Byteview，实现了 lru 包中定义的 Value 接口

type Byteview struct {
	b []byte
}

func (v Byteview) Len() int {
	return len(v.b)
}

func (v Byteview) String() string {
	return string(v.b)
}

func (v Byteview) ByteSlice() []byte {
	return CloneView(v.b)
}

func CloneView(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}