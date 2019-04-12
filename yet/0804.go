type Node struct {
    m string
    next *Node
}

var m_table []string = []string {".-","-...","-.-.","-..",".","..-.","--.","....",
                                 "..",".---","-.-",".-..","--","-.","---",".--.",
                                 "--.-",".-.","...","-","..-","...-",".--","-..-",
                                 "-.--","--.."}

func new(str string) *Node {
    return &Node{m: str, next: nil}
}

func c2m(c byte) string {
	return m_table[c - 'a']
}

func str2m(str string) string {
    res := ""
    for i := 0; i < len(str); i++ {
        res += c2m(str[i])
    }
    return res
}

func hash_m_str(m string) int {
    hash_val := 0

    // 0 <= dot_count <= 48
    // 0 <= bar_count <= 48
    dot_count := 0
    bar_count := 0
    for i := 0; i < len(m); i++ {
        if m[i] == '.' {
            dot_count++
            if i == 0 || i == len(m) - 1 {
                hash_val += 7
            }
        } else {
            bar_count++
            if i == 0 || i == len(m) - 1 {
                hash_val += 5
            }
        }
    }

    return hash_val + dot_count * bar_count / 3
}

func uniqueMorseRepresentations(words []string) int {
    hash_table := make([]bool, 1000)
    Node_table := make([]*Node, 1000)

    count := 0
    for _, word := range words {
        m := str2m(word)
        hash_val := hash_m_str(m)
        if (hash_table[hash_val]) {
            n := new(m)
            head := Node_table[hash_val]

            for {
                if n.m == head.m {
                    break
                }

                if head.next == nil {
                    count++
                    head.next = n
                }
                head = head.next
            }
        } else {
            hash_table[hash_val] = true
            count++
            Node_table[hash_val] = new(m)
        }
    }
    return count
}
