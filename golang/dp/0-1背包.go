package dp

func knapsack01(weight []int, value []int, bagWeight int) int {
	n := len(weight)
	if n == 0 || bagWeight == 0 {
		return 0
	}

	// 初始化 dp 数组：dp[i][j] 表示前 i 个物品在容量 j 时的最大价值
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
	}

	// 初始化第一个物品的情况
	for j := 0; j <= bagWeight; j++ {
		if j >= weight[0] {
			dp[0][j] = value[0]
		} else {
			dp[0][j] = 0
		}
	}

	// 动态规划填充表格
	for i := 1; i < n; i++ { // 遍历物品
		for j := 0; j <= bagWeight; j++ { // 遍历背包容量
			if j < weight[i] {
				dp[i][j] = dp[i-1][j] // 当前物品装不下，继承前一个状态
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}

	return dp[n-1][bagWeight]
}
