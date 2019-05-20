import "sort"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 1 {
		if amount%coins[0] == 0 {
			return amount / coins[0]
		}
		return -1
	}

	sort.Ints(coins)

	// [2, 3, 5] amount = 11

	//  0  1  2  3  4  5  6  7  8  9 10 11
	// [0,-1, 1,-1, 2,-1, 3,-1, 4,-1, 5,-1] <- 初期化

	// [0,-1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5] <- 3 = 3 * 1
	// [0,-1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4] <- 6 = 3 * 2
	// [0,-1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4] <- 9 = 3 * 3

	// [0,-1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3] <- 5 = 5 * 1
	// [0,-1, 1, 1, 2, 2, 2, 2, 2, 3, 2, 3] <- 10 = 5 * 2

	// index: コインの価値
	// 値: indexとなるためのコインの最小枚数
	memo := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		if i%coins[0] == 0 {
			memo[i] = i / coins[0]
		} else {
			memo[i] = -1
		}
	}

	for i := 1; i < len(coins); i++ {
		n := coins[i]
		for mul := 1; mul*n <= amount; mul++ {
			num := mul * n
			for j := mul * n; j <= amount; j++ {
				if memo[j-num] == -1 {
					memo[j] = memo[j]
				} else {
					if memo[j] == -1 {
						memo[j] = memo[j-num] + mul
					} else {
						memo[j] = Min(memo[j-num]+mul, memo[j])
					}
				}
			}
		}
	}

	return memo[amount]
}