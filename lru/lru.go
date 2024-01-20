package lru

import "container/list"

type kv struct {
	k, v int
}

type LRUCache struct {
	cap  int
	li   *list.List
	data map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:  capacity,
		li:   list.New(),
		data: make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.data[key]
	if !ok {
		return -1
	}
	this.li.MoveToFront(v)
	return v.Value.(kv).v
}

func (this *LRUCache) Put(key int, value int) {
	newCache := kv{
		k: key,
		v: value,
	}
	v, ok := this.data[key]
	if ok {
		v.Value = newCache
		this.li.MoveToFront(v)
		return
	}
	if this.isFull() {
		back := this.li.Back()
		this.li.Remove(back)
		delete(this.data, back.Value.(kv).k)
	}
	this.data[key] = this.li.PushFront(newCache)
}

func (this *LRUCache) isFull() bool {
	return len(this.data) == this.cap
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
