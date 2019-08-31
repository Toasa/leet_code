var chars_arr [26]int

func countCharacters(words []string, chars string) int {
	chars_arr = gen_string_arr(chars)

	// print_arr(chars, chars_arr)

	result := 0
	for _, word := range words {
		if isFormed(word, chars) {
			result += len(word)
		}
	}
	return result
}

func gen_string_arr(str string) [26]int {
	var arr [26]int
	for i := 0; i < 26; i++ {
		arr[i] = 0
	}

	for i := 0; i < len(str); i++ {
		index := str[i] - 'a'
		arr[index]++
	}
	return arr
}

func isFormed(word string, chars string) bool {
	if len(word) > len(chars) {
		return false
	}
	
	word_arr := gen_string_arr(word)

	// print_arr(word, word_arr)

	for i := 0; i < 26; i++ {
		if chars_arr[i] < word_arr[i] {
			return false
		}
	}
	return true
}

// for debug
func print_arr(word string, word_arr [26]int) {
	fmt.Println(word)
	for i := 0; i < 26; i++ {
		fmt.Printf("%c: %d   ", 'a' + i, word_arr[i])
	}
	fmt.Println()
}