package sort_map

type SortMap interface {
	Put(key Key, value Value)
	Get(key Key) (Value, bool)
	Delete(key Key)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
