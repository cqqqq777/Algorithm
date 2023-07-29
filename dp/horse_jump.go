package dp

// HorseJump 象棋跳马问题
func HorseJump(x, y, k int) int {
	return horseJumpRecur(x, y, k)
}

func horseJumpRecur(x, y, k int) int {
	if x < 0 || x > 8 || y < 0 || y > 9 {
		return 0
	}
	if k == 0 {
		if x == 0 && y == 0 {
			return 1
		}
		return 0
	}
	return horseJumpRecur(x-1, y+2, k-1) + horseJumpRecur(x+1, y+2, k-1) +
		horseJumpRecur(x+2, y+1, k-1) + horseJumpRecur(x+2, y-1, k-1) +
		horseJumpRecur(x+1, y-2, k-1) + horseJumpRecur(x-1, y-2, k-1) +
		horseJumpRecur(x-2, y-1, k-1) + horseJumpRecur(x-2, y+1, k-1)
}

func horseJumpDP(x, y, k int) int {
	dp := make([][][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([][]int, 9)
		for j := 0; j < 9; j++ {
			dp[i][j] = make([]int, 10)
		}
	}
	dp[0][0][0] = 1
	for i := 1; i < k+1; i++ {
		for j := 0; j < 9; j++ {
			for z := 0; z < 10; z++ {
				dp[i][j][z] += getLowLevelValue(dp, j-1, z-2, i-1)
				dp[i][j][z] += getLowLevelValue(dp, j-1, z+2, i-1)
				dp[i][j][z] += getLowLevelValue(dp, j+1, z-2, i-1)
				dp[i][j][z] += getLowLevelValue(dp, j+1, z+2, i-1)
				dp[i][j][z] += getLowLevelValue(dp, j-2, z-1, i-1)
				dp[i][j][z] += getLowLevelValue(dp, j-2, z+1, i-1)
				dp[i][j][z] += getLowLevelValue(dp, j+2, z-1, i-1)
				dp[i][j][z] += getLowLevelValue(dp, j+2, z+1, i-1)
			}
		}
	}
	return dp[k][x][y]
}

func getLowLevelValue(dp [][][]int, x, y, k int) int {
	if x < 0 || x > 8 || y < 0 || y > 9 {
		return 0
	}
	return dp[k][x][y]
}
