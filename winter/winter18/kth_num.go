package winter18

// 给定整数 n 和 k，返回  [1, n] 中字典序第 k 小的数字。
func findKthNumber(n int, k int) int {
	// 构造一个十叉前缀树树，前序遍历就是字典序
	cur := 1
	k--
	for k > 0 {
		steps := getSteps(cur, n)
		if k >= steps {
			// 跳出当前子树，往同层的右边走
			k -= steps
			cur++
		} else {
			// 进入子树的下一层
			k--
			cur *= 10
		}
	}
	return cur
}

func getSteps(cur, n int) int {
	start, end := cur, cur
	var steps int
	for start <= n {
		steps += min(end, n) - start + 1
		// 每一层的节点数 = 10 * 上一层的节点数
		start *= 10
		end = end*10 + 9
	}
	return steps
}
