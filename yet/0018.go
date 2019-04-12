func is_equal(n1, n2 []int) bool {
	for i := 0; i < 4; i++ {
		if n1[i] != n2[i] {
			return false
		}
	}
	return true
}

func fourSum(nums []int, target int) [][]int {

	res := [][]int{}
	sort.Ints(nums)
	l := len(nums)

	tmp := 0
	for h := 0; h < l - 3; h++ {
		tmp += nums[h]
		for i := h + 1; i < l - 2; i++ {
			tmp += nums[i]
			if tmp > target {
				tmp -= nums[i]
				continue
			}
			for j := i + 1; j < l - 1; j++ {
				tmp += nums[j]
				if tmp > target {
					tmp -= nums[j]
					continue
				}
				for k := j + 1; k < l; k++ {
					tmp += nums[k]
					if tmp == target {
						s := []int{nums[h], nums[i], nums[j], nums[k]}
						fmt.Println(s)
						if len(res) > 0 {
							if is_equal(s, res[len(res) - 1]) {
								tmp -= nums[k]
								continue
							}
						}
						res = append(res, s)
					}
					tmp -= nums[k]
				}
				tmp -= nums[j]
			}
			tmp -= nums[i]
		}
		tmp = 0
	}

	return res
}
