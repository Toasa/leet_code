import "sort"

func heightChecker(heights []int) int {
	sorted := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		sorted[i] = heights[i]
	}
	sort.Ints(sorted)

	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != sorted[i] {
			count++
		}
	}

	return count
}