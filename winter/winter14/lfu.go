package winter14

import (
	"container/list"
)

type Entry struct {
	key, val, frequency int
}

type LFUCache struct {
	cap, minFre int
	// data 记录数据
	data map[int]*list.Element
	// frequency 记录各个频次的链表
	frequency map[int]*list.List
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:       capacity,
		data:      make(map[int]*list.Element),
		frequency: make(map[int]*list.List),
	}
}

func (lfu *LFUCache) get(key int) *Entry {
	// get element
	element, ok := lfu.data[key]
	if !ok {
		return nil
	}
	entry := element.Value.(*Entry)

	// modify frequency map
	lfu.frequency[entry.frequency].Remove(element)

	entry.frequency++
	l, ok := lfu.frequency[entry.frequency]
	if !ok {
		l = list.New()
		element = l.PushFront(entry)
		lfu.frequency[entry.frequency] = l
	} else {
		element = l.PushFront(entry)
	}
	lfu.data[key] = element

	// update minFre
	l = lfu.frequency[lfu.minFre]
	if l.Len() == 0 {
		lfu.minFre = entry.frequency
	}

	return entry
}

func (lfu *LFUCache) Get(key int) int {
	if lfu.cap == 0 {
		return -1
	}
	if entry := lfu.get(key); entry != nil {
		return entry.val
	}
	return -1
}

func (lfu *LFUCache) Put(key int, value int) {
	if lfu.cap == 0 {
		return
	}
	entry := lfu.get(key)
	if entry == nil {
		// out of cap
		if len(lfu.data) == lfu.cap {
			minL := lfu.frequency[lfu.minFre]
			ele := minL.Back()
			minL.Remove(ele)
			delete(lfu.data, ele.Value.(*Entry).key)
		}
		l, ok := lfu.frequency[1]
		if !ok {
			l = list.New()
			lfu.frequency[1] = l
		}
		element := l.PushFront(&Entry{
			key:       key,
			val:       value,
			frequency: 1,
		})
		lfu.data[key] = element
		lfu.minFre = 1
	} else {
		entry.val = value
	}
}
