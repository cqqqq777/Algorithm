package greedy

import (
	"sort"
	"strings"
)

// LowestString 生成具有最小字段序的字符串
func LowestString(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return strings.Compare(strs[i], strs[j]) <= 0
	})
	return strings.Join(strs, "")
}
