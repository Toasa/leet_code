// Explanation:
// row 1: 0
// row 2: 01
// row 3: 0110
// row 4: 01101001
// row 5: 0110100110010110

// i >= 2のとき
// row i  : XY
// row i+1: XYYX

func rec(now, row int, s string) string {
	if now == row {
		return s
	}
	mid := len(s) / 2
	head := s[:mid]
	tail := s[mid:]
	s = rec(now+1, row, head+tail+tail+head)
	return s
}

// N = 30
// K = 434991989
// のときにruntime error

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	res := rec(2, N, "01")
    return int(res[K-1] - '0')
}