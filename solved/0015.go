import "sort"

func dup(res [][]int, tmp []int) bool {
	for _, s := range res {
		if s[0] == tmp[0] && s[1] == tmp[1] && s[2] == tmp[2] {
			return true
		}
	}
	return false
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 0 {
		return res
	}
	sort.Ints(nums)
	if 0 < nums[0] || nums[len(nums)-1] < 0 {
		return res
	}

	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if (nums[i] > nums[j]) || (nums[j] > -nums[i]-nums[j]) {
				continue
			}

			if m[nums[i]] == 1 {
				delete(m, nums[i])
			} else {
				m[nums[i]]--
			}

			if m[nums[j]] == 1 {
				delete(m, nums[j])
			} else {
				m[nums[j]]--
			}

			if _, ok := m[-nums[i]-nums[j]]; ok {
				tmp := []int{nums[i], nums[j], -nums[i] - nums[j]}
				//sort.Ints(tmp)

				if len(res) == 0 {
					res = append(res, tmp)
				} else {
					if !dup(res, tmp) {
						res = append(res, tmp)
					}
				}
			}

			m[nums[i]]++
			m[nums[j]]++
		}
	}

	return res
}