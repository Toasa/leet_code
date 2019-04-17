var d string

func solve(index int, res []string, l_comb string) []string {

	if index == len(d) {
		return res
	}

	switch d[index] {
	case '2':
		if index == len(d)-1 {
			res = append(res, l_comb+"a")
			res = append(res, l_comb+"b")
			res = append(res, l_comb+"c")
			return res
		} else {
			res = solve(index+1, res, l_comb+"a")
			res = solve(index+1, res, l_comb+"b")
			res = solve(index+1, res, l_comb+"c")
		}
	case '3':
		if index == len(d)-1 {
			res = append(res, l_comb+"d")
			res = append(res, l_comb+"e")
			res = append(res, l_comb+"f")
			return res
		} else {
			res = solve(index+1, res, l_comb+"d")
			res = solve(index+1, res, l_comb+"e")
			res = solve(index+1, res, l_comb+"f")
		}
	case '4':
		if index == len(d)-1 {
			res = append(res, l_comb+"g")
			res = append(res, l_comb+"h")
			res = append(res, l_comb+"i")
			return res
		} else {
			res = solve(index+1, res, l_comb+"g")
			res = solve(index+1, res, l_comb+"h")
			res = solve(index+1, res, l_comb+"i")
		}
	case '5':
		if index == len(d)-1 {
			res = append(res, l_comb+"j")
			res = append(res, l_comb+"k")
			res = append(res, l_comb+"l")
			return res
		} else {
			res = solve(index+1, res, l_comb+"j")
			res = solve(index+1, res, l_comb+"k")
			res = solve(index+1, res, l_comb+"l")
		}
	case '6':
		if index == len(d)-1 {
			res = append(res, l_comb+"m")
			res = append(res, l_comb+"n")
			res = append(res, l_comb+"o")
			return res
		} else {
			res = solve(index+1, res, l_comb+"m")
			res = solve(index+1, res, l_comb+"n")
			res = solve(index+1, res, l_comb+"o")
		}
	case '7':
		if index == len(d)-1 {
			res = append(res, l_comb+"p")
			res = append(res, l_comb+"q")
			res = append(res, l_comb+"r")
			res = append(res, l_comb+"s")
			return res
		} else {
			res = solve(index+1, res, l_comb+"p")
			res = solve(index+1, res, l_comb+"q")
			res = solve(index+1, res, l_comb+"r")
			res = solve(index+1, res, l_comb+"s")
		}
	case '8':
		if index == len(d)-1 {
			res = append(res, l_comb+"t")
			res = append(res, l_comb+"u")
			res = append(res, l_comb+"v")
			return res
		} else {
			res = solve(index+1, res, l_comb+"t")
			res = solve(index+1, res, l_comb+"u")
			res = solve(index+1, res, l_comb+"v")
		}
	case '9':
		if index == len(d)-1 {
			res = append(res, l_comb+"w")
			res = append(res, l_comb+"x")
			res = append(res, l_comb+"y")
			res = append(res, l_comb+"z")
			return res
		} else {
			res = solve(index+1, res, l_comb+"w")
			res = solve(index+1, res, l_comb+"x")
			res = solve(index+1, res, l_comb+"y")
			res = solve(index+1, res, l_comb+"z")
		}
	}

	return res
}

// 2->abc, 3->def, 4->ghi, 5->jkl, 6->mno, 7->pqrs, 8->tuv, 9->wxyz
func letterCombinations(digits string) []string {
	d = digits
	res := []string{}
	res = solve(0, res, "")
	return res
}