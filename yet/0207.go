const (
    NONE = iota
    HEAD
    TAIL
    H_T
)

func changeStatus(prev, next uint8) uint8 {
    if prev == next {
        return next
    }
    if prev == NONE {
        return next
    }
    if prev == HEAD {
        return H_T
    }
    if prev == TAIL {
        return H_T
    }

    return 10
}

var white_set []int
var gray_set []int
var black_set []int

func move(cur int, src, dst string) {
    a
}

func dfs(cur int, white_set, gray_set, black_set []int{}) bool {
    move(cur, white_set, gray_set)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
    for i := 0; i < numCourses; i++ {
        white_set = append(white_set, i)
    }

    for len(white_set) > 0 {
        cur := white_set[0];
        white_set = white_set[1:]
        if dfs(cur, white_set, gray_set, black_set) {
            return false
        }
    }
    return true


    fromHead := make([][]int, numCourses)
    toTail := make([][]int, numCourses)

    // head and tail array
    // 0: none, 1: head, 2: tail, 3:head and tail
    state := make([]uint8, numCourses)
    head_num := 0

    for _, pair := range prerequisites {
        h := pair[1]
        t := pair[0]

        before := state[h]
        state[h] = changeStatus(state[h], HEAD)
        if state[h] == HEAD && before != HEAD {
            head_num++
        } else if state[h] != HEAD && before == HEAD {
            head_num--
        }

        state[t] = changeStatus(state[t], TAIL)

        fromHead[h] = append(fromHead[h], t)
        toTail[t] = append(toTail[t], h)
    }

    if head_num == 0 {
        return false
    }

    node_count := numCourses
    for head_num > 0 {
        if head_num == 0 && node_count > 0 {
            return false
        }

        var h_node uint8
        for i := 0; i < numCourses; i++ {
            if state[i] == HEAD {
                h_node = i
                break
            }
        }
        state[h_node] = NONE
        head_num--
        node_count--
    }

    return true

    fmt.Println(head_num)
    fmt.Println(state)
    fmt.Println(fromHead)

    return true
}
