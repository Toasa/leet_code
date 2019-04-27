func find(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return
		}
	}

	mid := len(nums) / 2
	if nums[mid] == target {
		return mid
	}

	if target < nums[mid] {
		return find(nums[:mid], target)
	}

	if nums[mid] < target {
		return find(nums[mid:], target)
	}
}

func search(nums []int, target int) int {
	half := len(nums) / 2

	var pivot int

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	if len(nums) == 2 {
		if nums[0] == target {
			return 0
		} else if nums[1] == target {
			return 1
		}
		return -1
	}

	for i := 0; 0 <= half-i; i++ {
		if nums[half-i-1] > nums[half-i] {
			pivot = i
			break
		}
		if nums[half+i-1] > nums[half+i] {
			pivot = i
			break
		}
	}

	head := nums[:pivot]
	tail := nums[pivot:]

	var res int
	if head[0] < target && target < head[len(head)-1] {
		res = find(head, target)
	}
	if tail[0] < target && target < tail[len(tail)-1] {
		res = find(tail, target) + half
	}

	return res
}