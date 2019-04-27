import "strings"

var N int

// opc, cpc: open / closed parenthesis count,
func solve(opc, cpc int, s string, res []string) []string {
	// if opc < cpc {
	// 	fmt.Println("Error")
	// 	return nil
	// }

	if opc == N {
		res = append(res, s+strings.Repeat(")", opc-cpc))
		return res
	}
	res = solve(opc+1, cpc, s+"(", res)
	if cpc < opc {
		res = solve(opc, cpc+1, s+")", res)
	}
	return res
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	N = n
	res := []string{}
	res = solve(1, 0, "(", res)
	return res
}