// apple => apple + ma
// goat => oat + g + ma

// I apple goat =>
// Imaa applemaaa oatgmaaaa

var head int = 0

func is_vowel(c byte) bool {
    switch c {
    case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
        return true
    default:
        return false
    }
}

func is_space(c byte) bool {
    return c == ' '
}

func skip(S string) {
    for is_space(S[head]) {
        head++
        if head >= len(S) {
            return
        }
    }
}

func read_word(S string) int {
    l := 1
    if head + l >= len(S) {
        return l
    }
    for !is_space(S[head + l]) {
        l++
        if head + l >= len(S) {
            return l
        }
    }
    return l
}

func convert(S string) string {
    if is_vowel(S[0]) {
        return S + "ma"
    } else {
        return S[1:] + string(S[0]) + "ma"
    }
}

func toGoatLatin(S string) string {
    res := ""

    for a_count := 1; head < len(S); a_count++ {
        skip(S)
        if head >= len(S) {
            return res
        }

        word_len := read_word(S)
        res += convert(S[head : head + word_len]) + strings.Repeat("a", a_count) + " "
        head += word_len
    }

    // fmt.Println(res)
    // fmt.Println(len(res) - 1)

    return res[:len(res) - 1]
}
