import "fmt"

var A [][]int
var Nums []int
var l int

func p(tmp []int, memo []bool) {
	if len(tmp) >= l {
		A = append(A, tmp)
		// fmt.Println(tmp)
		return
	}

	for i := 0; i < len(memo); i++ {
		fmt.Println(i, ":", tmp)
		if !memo[i] {
			memo[i] = true
			tmp = append(tmp, Nums[i])

			p(tmp, memo)

			memo[i] = false
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func permute(nums []int) [][]int {
	A = make([][]int, 0)
	Nums = nums
	l = len(nums)

	tmp := []int{}
	memo := make([]bool, l)
	p(tmp, memo)

	return A
}