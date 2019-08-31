func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	
	for i := 0; i < len(nums); i++ {
		freqMap[nums[i]]++
	}

	
}