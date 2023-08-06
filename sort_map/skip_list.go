/*
Package sort_map provides a map with Skip List.
TODO: implement Delete method.
*/

package sort_map

import (
	"math"
	"math/rand"
)

type SkipListNode struct {
	key   Key
	value Value
	next  []*SkipListNode
}

type SkipList struct {
	head     *SkipListNode
	maxLevel int
}

func NewSkipList() *SkipList {
	return &SkipList{head: &SkipListNode{next: make([]*SkipListNode, 1)}}
}

func (sl *SkipList) Get(key Key) (Value, bool) {
	mostRight := sl.findMostRightNode(key, sl.maxLevel)
	if mostRight != nil && mostRight.key == key {
		return mostRight.value, true
	}
	return 0, false
}

func (sl *SkipList) Put(key Key, value Value) {
	level := 0
	// 25% chance to increase level
	for rand.Intn(5) == 0 {
		level++
	}
	if sl.maxLevel < level {
		sl.head.next = append(sl.head.next, make([]*SkipListNode, level-sl.maxLevel)...)
		sl.maxLevel = level
	}
	mostRight := sl.findMostRightNode(key, sl.maxLevel)
	if mostRight != nil && mostRight.key == key {
		mostRight.value = value
		return
	}
	// insert new node
	newNode := &SkipListNode{key: key, value: value, next: make([]*SkipListNode, level+1)}
	nowLevel := int(math.Min(float64(len(mostRight.next)), float64(level+1)))
	for i := 0; i < nowLevel; i++ {
		newNode.next[i] = mostRight.next[i]
		mostRight.next[i] = newNode
	}
	mostRight = sl.findMostRightInLevel(key, level)
	for ; nowLevel <= level; nowLevel++ {
		newNode.next[nowLevel] = mostRight.next[nowLevel]
		mostRight.next[nowLevel] = newNode
	}
}

func (sl *SkipList) Delete(key Key) {
	mostRight := sl.head
	level := sl.maxLevel
	for level >= 0 {
		for mostRight.next[level] != nil && mostRight.next[level].key < key {
			mostRight = mostRight.next[level]
		}
		if mostRight.next[level] != nil && mostRight.next[level].key == key {
			mostRight.next[level] = mostRight.next[level].next[level]
			level--
		} else {
			level--
		}
	}
}

func (sl *SkipList) findMostRightNode(key Key, level int) *SkipListNode {
	cur := sl.head
	for level >= 0 {
		for cur.next[level] != nil && cur.next[level].key <= key {
			cur = cur.next[level]
		}
		level--
	}
	return cur
}

func (sl *SkipList) findMostRightInLevel(key Key, level int) *SkipListNode {
	cur := sl.head
	for cur.next[level] != nil && cur.next[level].key < key {
		cur = cur.next[level]
	}
	return cur
}
