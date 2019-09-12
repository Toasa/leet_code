func alignString(a, b string) (string, string) {
	if (len(a) < len(b)) {
		a = strings.Repeat("0", len(b) - len(a)) + a
	} else {
		b = strings.Repeat("0", len(a) - len(b)) + b
	}
	return a, b
}

func addBinary(a string, b string) string {
	a, b = alignString(a, b)
	
	res := ""
	carryFlag := false

	for i := len(a) - 1; 0 <= i; i-- {
		if a[i] == '0' && b[i] == '0' {
			if carryFlag {
				res = "1" + res
				carryFlag = false
			} else {
				res = "0" + res
			}			
		} else if (a[i] == '1' && b[i] == '0') || (a[i] == '0' && b[i] == '1') {
			if carryFlag {
				res = "0" + res
			} else {
				res = "1" + res
			}
		} else if a[i] == '1' && b[i] == '1' {
			if carryFlag {
				res = "1" + res
			} else {
				res = "0" + res
			}
			carryFlag = true
		}
	}

	if carryFlag {
		res = "1" + res
	}

	return res
}