package winter11

// 给你一棵二叉树的根节点 root ，返回树的 最大宽度 。
// 树的 最大宽度 是所有层中最大的 宽度 。
// 每一层的 宽度 被定义为该层最左和最右的非空节点（即，两个端点）之间的长度。
// 将这个二叉树视作与满二叉树结构相同，两端点间会出现一些延伸到这一层的 null 节点，这些 null 节点也计入长度。
// 题目数据保证答案将会在  32 位 带符号整数范围内。
type node struct {
	val *TreeNode
	idx int
}

// 层序遍历每次记一下索引即可
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*node, 0)
	var idx, res int
	queue = append(queue, &node{
		val: root,
		idx: idx,
	})
	for len(queue) != 0 {
		sz := len(queue)
		level := make([]*node, 0)
		for i := 0; i < sz; i++ {
			n := queue[0]
			queue = queue[1:]
			idx = n.idx
			level = append(level, n)
			if n.val.Left != nil {
				queue = append(queue, &node{
					val: n.val.Left,
					idx: idx*2 + 1,
				})
			}
			if n.val.Right != nil {
				queue = append(queue, &node{
					val: n.val.Right,
					idx: idx*2 + 2,
				})
			}
		}
		res = max(level[len(level)-1].idx-level[0].idx, res)
	}
	return res
}
