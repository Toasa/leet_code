func isSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}

	t_i := -1
	for i := 0; i < len(s); i++ {
		for j := t_i + 1; ; j++ {
			if j >= len(t) {
				return false
			}

			if s[i] == t[j] {
				t_i = j
				break
			}
		}
	}
	return true
}