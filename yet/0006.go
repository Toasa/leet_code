func convert(s string, numRows int) string {
    periodic_strs := []string{}
    one_period := 2 * numRows - 2

    head := 0

    for head + one_period <= len(s) {
        periodic_strs = append(periodic_strs, s[head : head + one_period])
        head += one_period
    }

    if head < len(s) {
        periodic_strs = append(periodic_strs, s[head : len(s)])
    }

    fmt.Println(periodic_strs)
}
