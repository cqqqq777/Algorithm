package concurrent_map

import (
	"errors"
	"sync"
	"time"
)

var ErrTimeExpire = errors.New("time out")

// MyConcurrentMap 面向并发场景
// 提供 Put 与 Get 两个方法，时间复杂度 O(1)
// 查询时若 key 不存在则阻塞并等待 maxWaitingTime 时间，若在该时间内有写入该 key 的操作，返回 value，超时则返回超时错误
// 不能有死锁和 panic 风险
type MyConcurrentMap struct {
	sync.Locker
	mp      map[int]int
	keyToCh map[int]chan struct{}
}

func NewConcurrentMap(l sync.Locker) *MyConcurrentMap {
	return &MyConcurrentMap{
		l,
		make(map[int]int),
		make(map[int]chan struct{}),
	}
}

func (m *MyConcurrentMap) Put(k, v int) {
	m.Lock()
	defer m.Unlock()
	m.mp[k] = v
	// 注意这里不能使用发送信号的方式解锁，这样只能解锁一个读阻塞的协程
	// 根据 channel 的底层，直接关闭 channel 可以解锁所有阻塞的协程
	if ch, ok := m.keyToCh[k]; ok {
		close(ch)
		delete(m.keyToCh, k)
	}
}

func (m *MyConcurrentMap) Get(k int, maxWaitingTime time.Duration) (int, error) {
	m.Lock()
	if v, ok := m.mp[k]; ok {
		m.Unlock()
		return v, nil
	}
	// 先判断此 key 先前是否被操作
	ch, ok := m.keyToCh[k]
	if !ok {
		m.keyToCh[k] = make(chan struct{})
	}
	m.Unlock()
	timer := time.NewTimer(maxWaitingTime)
	select {
	case <-timer.C:
		return 0, ErrTimeExpire
	case <-ch:
		m.Lock()
		v := m.mp[k]
		m.Unlock()
		return v, nil
	}
}
