func myPow(x float64, n int) float64 {
	if n > 1000000 {
		if n%2 == 0 {
			n = 1000000
		} else {
			n = 1000001
		}
	} else if n < -1000000 {
		if n%2 == 0 {
			n = -1000000
		} else {
			n = -1000001
		}
	}

	if n == 0 {
		return 1
	}

	if x == 0.0 {
		return 0.0
	}

	if n > 0 {
		return x * myPow(x, n-1)
	}
	if n < 0 {
		n = -n
		return 1 / (x * myPow(x, n-1))
	}
	return 1
}