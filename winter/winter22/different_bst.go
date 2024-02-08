package winter22

// 给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？
// 返回满足题意的二叉搜索树的种数。
func numTrees(n int) int {
	// 节点数一定，能构成的 bst 数量就一定
	// n 个节点构成的 bst 数就等于左子树的 bst 数 * 右子树的 bst 数
	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 1, 1, 2
	for i := 3; i <= n; i++ {
		for j := 0; j < i; j++ {
			// 枚举左右子树的节点分配情况
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
