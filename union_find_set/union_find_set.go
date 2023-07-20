package union_find_set

type Element struct {
	Value int
}

type UnionFindSet struct {
	FatherMap map[*Element]*Element
	SizeMap   map[*Element]int
}

// NewUnionFindSet init a union find set
func NewUnionFindSet(elements ...*Element) *UnionFindSet {
	ufs := &UnionFindSet{FatherMap: make(map[*Element]*Element), SizeMap: make(map[*Element]int)}
	for _, v := range elements {
		ufs.FatherMap[v] = v
		ufs.SizeMap[v] = 1
	}
	return ufs
}

// IsSameSet check the a and b whether reside in the same set
func (u *UnionFindSet) IsSameSet(a, b *Element) bool {
	if a == b {
		return true
	}
	headA, ok1 := u.FindHead(a)
	headB, ok2 := u.FindHead(b)
	if ok1 && ok2 && headA == headB {
		return true
	}
	return false
}

// Union merge the sets where a and b reside
func (u *UnionFindSet) Union(a, b *Element) {
	if a == b {
		return
	}
	headA, ok1 := u.FindHead(a)
	headB, ok2 := u.FindHead(b)
	if ok1 && ok2 && headA != headB {
		bigger, smaller := headA, headB
		if u.SizeMap[headA] < u.SizeMap[headB] {
			bigger = headB
			smaller = headA
		}
		u.FatherMap[smaller] = bigger
		u.SizeMap[bigger] += u.SizeMap[smaller]
		delete(u.SizeMap, smaller)
	}
}

// FindHead
// if element != nil, return the root node and true, otherwise nil and false
// if element not in the union find set, create a set for it
func (u *UnionFindSet) FindHead(element *Element) (*Element, bool) {
	if element == nil {
		return nil, false
	}
	if _, ok := u.FatherMap[element]; !ok {
		u.FatherMap[element] = element
		u.SizeMap[element] = 1
		return element, true
	}
	path := make([]*Element, 0)
	for i := 0; u.FatherMap[element] != element; i++ {
		path = append(path, element)
		element = u.FatherMap[element]
	}
	for i := len(path) - 1; i >= 0; i-- {
		u.FatherMap[path[i]] = element
	}
	return element, true
}
