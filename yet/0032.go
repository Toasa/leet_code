import "fmt"

var max int
var str string

// parenthesis count
var pc int

func solve(i int) {
	tmp := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			pc++
		} else {
			pc--
			if pc >= 0 {
				tmp++
			}
		}

		if tmp > max {
			max = tmp
		}

		if pc < 0 {
			tmp = 0
			pc = 0
		}
		fmt.Println("index:", i, "pc:", pc, "max:", max)
	}
}

func longestValidParentheses(s string) int {
	str = s
	max = 0
	pc = 0

	var i int

	solve(i)

	return 2 * max
}