import "fmt"

// board[i][j]において、第i行目に数字j+1が挿入済みであるか否か
var row_done [9][9]bool
var col_done [9][9]bool

// 0 1 2
// 3 4 5
// 6 7 8
var cell_done [9][9]bool

var cell_option [][][]int

func get_cell_index(i, j int) int {
	if 0 <= i && i <= 2 {
		if 0 <= j && j <= 2 {
			return 0
		} else if 3 <= j && j <= 5 {
			return 1
		} else {
			return 2
		}
	} else if 3 <= i && i <= 5 {
		if 0 <= j && j <= 2 {
			return 3
		} else if 3 <= j && j <= 5 {
			return 4
		} else {
			return 5
		}
	} else {
		if 0 <= j && j <= 2 {
			return 6
		} else if 3 <= j && j <= 5 {
			return 7
		} else {
			return 8
		}
	}
}

func get_option(row, col int) []int {
	res := []int{}
	for num := 0; num < 9; num++ {
		if row_done[row][num] {
			continue
		}
		if col_done[col][num] {
			continue
		}
		if cell_done[get_cell_index(row, col)][num] {
			continue
		}
		res = append(res, num)
	}
	return res
}

func update(b [][]byte, i, j, num int) {
	b[i][j] = byte(num + 49)
	// b, row_done, col_done, cell_doneの更新
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := b[i][j]
			if c != '.' {
				row_done[i][c-'1'] = true
				col_done[j][c-'1'] = true
				cell_done[get_cell_index(i, j)][c-'1'] = true
			}
		}
	}
}

func fill_fixed_cells(b [][]byte) {
	var i, j int
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			if b[i][j] == '.' {
				// cell_option[i][j]に候補となる数字を挿れる
				cell_option[i][j] = get_option(i, j)
				if len(cell_option[i][j]) == 1 {
					update(b, i, j, cell_option[i][j][0])
					cell_option[i][j] = get_option(i, j)
					i = 0
					j = 0
					break
				}
			}
		}
	}
}

func initi(b [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := b[i][j]
			if c != '.' {
				row_done[i][c-'1'] = true
				col_done[j][c-'1'] = true
				cell_done[get_cell_index(i, j)][c-'1'] = true
			}
		}
	}

	cell_option = make([][][]int, 9)
	for i := 0; i < 9; i++ {
		cell_option[i] = make([][]int, 9)
	}

	fill_fixed_cells(b)
}

func is_complete(b [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == '.' {
				return false
			}
		}
	}
	return true
}

func is_conflict(b [][]byte, i, j int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

		}
	}
	return true
}

func solve(b [][]byte) bool {
	if is_complete(b) {
		return true
	}

	// i, j, numを一組選び、挿入する
	tmp_i, tmp_j, tmp_num := get_certain_pair()
	b[tmp_i][tmp_j] = tmp_num
	if is_conflict(b, tmp_i, tmp_j) {
		return false
	}

	solve(b)
}

// boardには1 - 9の数字を
// その他の配列には0 - 8の数字を入れる
func solveSudoku(board [][]byte) {
	initi(board)

	solve(board)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}