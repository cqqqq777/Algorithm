package hash

import "math/rand"

type RandomPool struct {
	keyValue map[string]int
	valueKey map[int]string
}

func NewRandomPool() *RandomPool {
	return &RandomPool{keyValue: make(map[string]int), valueKey: make(map[int]string)}
}

func (r *RandomPool) Insert(key string) {
	if r == nil || r.keyValue == nil || r.valueKey == nil {
		panic("nil pointer")
	}
	value := len(r.valueKey)
	r.keyValue[key] = value
	r.valueKey[value] = key
}

func (r *RandomPool) GetRandom() string {
	if len(r.keyValue) == 0 {
		return ""
	}
	random := rand.Intn(len(r.valueKey))
	return r.valueKey[random]
}

func (r *RandomPool) Delete(key string) {
	if _, ok := r.keyValue[key]; !ok {
		return
	}
	size := len(r.valueKey)
	index := r.keyValue[key]
	last := r.valueKey[size]
	r.valueKey[index] = r.valueKey[size]
	r.keyValue[last] = index
	delete(r.valueKey, size)
	delete(r.keyValue, key)
}
