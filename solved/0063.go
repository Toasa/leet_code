func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// メモ化すればよい
	// - +
	// + ?   <- 上と左の+から?を求める

	// 初期化も必要
	// ++++++++++++++++
	// +---------------
	// +---------------
	// +---------------

	row := len(obstacleGrid)
	col := len(obstacleGrid[0])

	memo := make([][]int, row)
	for i := 0; i < row; i++ {
		memo[i] = make([]int, col)
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	} else {
		memo[0][0] = 1
	}

	var obsFlg = false
	// memoの初期化
	for i := 1; i < row; i++ {
		if obsFlg {
			memo[i][0] = 0
			continue
		}
		if obstacleGrid[i][0] == 1 {
			obsFlg = true
			memo[i][0] = 0
		}
		if !obsFlg && obstacleGrid[i][0] == 0 {
			memo[i][0] = 1
		}
	}
	obsFlg = false
	for i := 1; i < col; i++ {
		if obsFlg {
			memo[0][i] = 0
			continue
		}
		if obstacleGrid[0][i] == 1 {
			obsFlg = true
			memo[0][i] = 0
		}
		if !obsFlg && obstacleGrid[0][i] == 0 {
			memo[0][i] = 1
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] == 1 {
				memo[i][j] = 0
			} else {
				memo[i][j] = memo[i-1][j] + memo[i][j-1]
			}
		}
	}
	return memo[row-1][col-1]
}