func firstUniqChar(s string) int {
	memo := [26]int{}
	for i := 0; i < len(s); i++ {
		memo[s[i] - 'a']++
	}

	for i := 0; i < len(s); i++ {
		if memo[s[i] - 'a'] == 1 {
			return i;
		}
	}

	return -1
}