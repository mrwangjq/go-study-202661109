package main

import "fmt"

type HashMap[K comparable, V any] struct {
	data map[K]V
}

// 构造函数
func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{
		data: make(map[K]V),
	}
}

// Set 设置键值
func (h *HashMap[K, V]) Set(key K, val V) {
	h.data[key] = val
}

// Get 获取
func (h *HashMap[K, V]) Get(key K) (V, bool) {
	v, ok := h.data[key]
	return v, ok
}

// Delete 删除
func (h *HashMap[K, V]) Delete(key K) {
	delete(h.data, key)
}

func main() {
	hm := NewHashMap[string, int]()
	hm.Set("age", 22)
	val, ok := hm.Get("age")
	if ok {
		fmt.Println(val)
	}
}
