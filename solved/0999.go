func print_board(board [][]byte) {
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            fmt.Printf("%c ", board[i][j])
        }
        fmt.Printf("\n")
    }
}

func get_R_position(board [][]byte) int {
  for row := 0; row < 8; row++ {
    for col := 0; col < 8; col++ {
      if board[row][col] == 'R' {
        return 8 * row + col
      }
    }
  }
  return 64
}

func get_ans(b [][]byte, pos_R int) int {

  row_R := pos_R / 8
  col_R := pos_R % 8

  // NWSE
  checked := []bool{false, false, false, false}

  res := 0

  for i := 0; i < 8; i++ {
    // N
    if !checked[0] {
      if row_R - i < 0 || b[row_R - i][col_R] == 'B' {
        checked[0] = true
      } else if b[row_R - i][col_R] == 'p' {
        checked[0] = true
        res++
      }
    }
    // W
    if !checked[1] {
      if col_R - i < 0 || b[row_R][col_R - i] == 'B' {
        checked[1] = true
      } else if b[row_R][col_R - i] == 'p' {
        checked[1] = true
        res++
      }
    }
    // S
    if !checked[2] {
      if row_R + i > 7 || b[row_R + i][col_R] == 'B' {
        checked[2] = true
      } else if b[row_R + i][col_R] == 'p' {
        checked[2] = true
        res++
      }
    }
    // E
    if !checked[3] {
      if col_R + i > 7 || b[row_R][col_R + i] == 'B' {
        checked[3] = true
      } else if b[row_R][col_R + i] == 'p' {
        checked[3] = true
        res++
      }
    }
  }

  return res
}

func numRookCaptures(board [][]byte) int {
    // print_board(board)
    pos_R := get_R_position(board)
    return get_ans(board, pos_R)
}
