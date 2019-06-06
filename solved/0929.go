func opt(mail string) string {
	res := ""

	i := 0
	for ; mail[i] != '@'; i++ {
		if mail[i] == '.' {
			continue
		} else if mail[i] == '+' {
			for mail[i+1] != '@' {
				i++
			}
			continue
		} else {
			res += string(mail[i])
		}			
	}

	res += mail[i:]
    
    fmt.Println(res)
	return res
}

func numUniqueEmails(emails []string) int {
	memo := map[string]byte{}
	count := 0
	
	for _, email := range emails {
		email = opt(email)
		_, ok := memo[email]
		if !ok {
			memo[email] = 1
			count++
		}
	}

	return count
}