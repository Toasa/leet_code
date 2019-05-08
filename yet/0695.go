import "fmt"

// input = [[1,1,0,1,1],[1,0,0,0,0],[0,0,0,0,1],[1,1,0,1,1]]
// でWA
// だが、"Run code"では正しい結果が出ている。何故だろう？

var zero_queue [][]int
var land_queue [][]int

// 0: 未, 1: 訪問待ち, 2: 訪問済
var checked [50][50]uint8
var row_l int
var col_l int

func print_checked() {
	for i := 0; i < row_l; i++ {
		for j := 0; j < col_l; j++ {
			fmt.Printf("%d ", checked[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func adj_check(grid_val, r, c int) {
	checked[r][c] = 1
	if grid_val == 1 {
		land_queue = append(land_queue, []int{r, c})
	} else {
		zero_queue = append(zero_queue, []int{r, c})
	}
}

func maxAreaOfIsland(grid [][]int) int {
	row_l = len(grid)
	col_l = len(grid[0])

	zero_queue = make([][]int, 0)
	land_queue = make([][]int, 0)
	cur_size := 0
	max_size := 0
	visit_count := 0

	if grid[0][0] == 1 {
		land_queue = append(land_queue, []int{0, 0})
	} else {
		zero_queue = append(zero_queue, []int{0, 0})
	}

	for visit_count < row_l*col_l {
		var cell []int
		if len(land_queue) > 0 {
			cell = land_queue[0]
			land_queue = land_queue[1:]
			cur_size++
		} else {
			if cur_size > max_size {
				max_size = cur_size
			}
			cur_size = 0
			if len(zero_queue) == 0 {
				return max_size
			}
			cell = zero_queue[0]
			zero_queue = zero_queue[1:]
		}

		row_i := cell[0]
		col_i := cell[1]

		checked[row_i][col_i] = 2
		visit_count++

		if visit_count == row_l*col_l {
			return max_size
		}

		if row_i > 0 {
			if checked[row_i-1][col_i] == 0 {
				adj_check(grid[row_i-1][col_i], row_i-1, col_i)
			}
		}
		if row_i+1 < row_l {
			if checked[row_i+1][col_i] == 0 {
				adj_check(grid[row_i+1][col_i], row_i+1, col_i)
			}
		}
		if col_i > 0 {
			if checked[row_i][col_i-1] == 0 {
				adj_check(grid[row_i][col_i-1], row_i, col_i-1)
			}
		}
		if col_i+1 < col_l {
			if checked[row_i][col_i+1] == 0 {
				adj_check(grid[row_i][col_i+1], row_i, col_i+1)
			}
		}

		// print_checked()
	}
	return max_size

	// while visit_count < row_l * col_l:
	//     len(land_queue) > 0 の場合
	//         land_queueからdequeし、それをcellとおく。
	//         また、cur_size++する。そしてcheckedを更新する。
	//     len(land_queue) == 0 の場合
	//         cur_sizeがmax_sizeより大きい場合更新する
	//         cur_size = 0に戻す
	//         len(zero_queue) == 0の場合
	//             return max_size
	//         else:
	//             zero_queueからdequeし、それをcellとおく。
	//             checkedを更新する。
	//
	//     visit_count += 1
	//
	//     cellの上下左右のマスを調べる
	//     隣接するマスがchecked == falseであり、かつ
	//	      -> 1ならland_queueにenqueする
	//        -> 0ならzero_queueにenqueする
}