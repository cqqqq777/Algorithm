package day06

// PrintPos
// 已知一棵二叉树中没有重复节点，并且给定了这棵树的中序遍历数组和先序遍历数组，返回后序遍历数组。
// 比如给定：
// int[] pre = { 1, 2, 4, 5, 3, 6, 7 };
// int[] in = { 4, 2, 5, 1, 6, 3, 7 };
// 返回：
// {4,5,2,6,7,3,1}
func PrintPos(pre, in []int) []int {
	size := len(pre)
	pos := make([]int, size)
	set(pre, in, pos, 0, size-1, 0, size-1, 0, size-1)
	return pos
}

func set(pre, in, pos []int, preI, preJ, inI, inJ, posI, posJ int) {
	if preI > preJ {
		return
	}
	if preI == preJ {
		pos[posJ] = pre[preI]
		return
	}
	pos[posJ] = pre[preI]
	find := inI
	for ; find <= inJ; find++ {
		if pre[preI] == in[find] {
			break
		}
	}
	set(pre, in, pos, preI+1, preI+find-inI, inI, find-1, posI, posI+find-inI-1)
	set(pre, in, pos, preI+find-inI+1, preJ, find+1, inJ, posI+find-inI, posJ-1)
}
