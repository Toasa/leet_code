func longestPalindrome(s string) string {
    l := len(s)
    dp := make([][]bool, l)
    for i := 0; i < l; i++ {
        dp[i] = make([]bool, l)
    }

    memo := [2]int{0, 0}

    if l == 0 {
        return ""
    }
    if l == 1 {
        return string(s[0]);
    }
    if l == 2 {
        if s[0] == s[1] {
            return s
        } else {
            return string(s[0])
        }
    }

    // initialize
    for i := 0; i < l; i++ {
        dp[i][i] = true
    }
    for i := 0; i < l - 1; i++ {
        if s[i] == s[i + 1] {
            memo = [2]int{i, i+1}
            dp[i][i + 1] = true
        } else {
            dp[i][i + 1] = false
        }
    }

    for rng := 2; rng < l; rng++ {
        for head := 0; head < l - rng; head++ {
            tail := head + rng

            if s[head] == s[tail] {
                if dp[head + 1][tail - 1] {
                    dp[head][tail] = true
                    memo = [2]int{head, tail}
                } else {
                    dp[head][tail] = false
                }
            } else {
                dp[head][tail] = false
            }
        }
    }

    // for i := 0; i < l; i++ {
    //     for j := 0; j < l; j++ {
    //         fmt.Printf("%d ", dp[i][j])
    //     }
    //     fmt.Println()
    // }

    return s[memo[0] : memo[1]+1]
}
