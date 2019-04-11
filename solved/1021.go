type stack struct {
    size int
    cont []byte
}

func new_stack() *stack {
    c := []byte{}
    return &stack{size: 0, cont: c}
}

func (s *stack) push(b byte) {
    s.cont = append(s.cont, b)
    s.size++
}

func (s *stack) pop() byte {
    b := s.cont[s.size - 1]
    s.cont = s.cont[ : s.size - 1]
    s.size--
    return b
}

func removeOuterParentheses(S string) string {
    s := new_stack()

    remove_indices := []int{}

    for i := 0; i < len(S); i++ {
        if S[i] == '(' {
            if s.size == 0 {
                remove_indices = append(remove_indices, i)
            }
            s.push(S[i])
        }
        if S[i] == ')' {
            s.pop()
            if s.size == 0 {
                remove_indices = append(remove_indices, i)
            }
        }
    }

    for i := 0; i < len(remove_indices); i++ {
        if remove_indices[i] == 0 {
            S = S[1:]
        } else if remove_indices[i] == len(S) - 1 {
            S = S[:len(S) - 1 - i]
        } else {
            S = S[:remove_indices[i] - i] + S[remove_indices[i] - i + 1 :]
        }
    }

    return S
}
