func faq(n int) int {
	if n == 0 {
		return 1
	}
	return n * faq(n-1)
}

func C(n, k int) int {
	if n == k {
		return 1
	}

	n_tmp := 1
	for i := 0; i < k; i++ {
		n_tmp *= n
		n--
	}
	return n_tmp / faq(k)
}

func uniquePaths(m int, n int) int {
	if m == 1 && n == 1 {
		return 1
	}

	m--
	n--
	if n < m {
		return C(m+n, n)
	}
	return C(m+n, m)
}