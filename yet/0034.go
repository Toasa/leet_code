var t int
var Nums []int

// [1, 2, 3, 4, 5, 6] => mid is 3
// [1, 2, 3, 4, 5] => mid is 3

// [3, 4]

// head, tail is index of currently searching
func find(head, tail int) bool, int {
	if tail - head == 1 {
		if Nums[tail] == t {
			return true, tail
		}
		if Nums[head] == t {
			return true, head
		}
		return false, nil
	}

	mid = (head + tail) / 2
	if Nums[mid] == t {
		return true, mid
	} else if Nums[mid] < t {
		return find(mid, tail)
	} else {
		return find(head, mid)
	}
}

func searchRange(nums []int, target int) []int {
	t = target
	Nums = nums
	res, index = find(0, len(Nums) - 1)

	if !res {
		return []int{-1, -1}
	} else {
		head := index
		tail := index
		for int i = 1; 0 <= index - i || index + i < len(Nums) {
			if 0 <= index-i {
				if Nums[index-i] == t {
					head--
				}
			}
			if index+i < len(Nums) {
				if Nums[index-i] == t {
					tail++
				}
			}
		}
		return []int{head, tail}
	}
}