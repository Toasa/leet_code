func fourSum(nums []int, target int) [][]int {
    ans := [][]int{}
    L := len(nums)

    for h := 0; h < L - 3; h++ {
        for i := h + 1; h < L - 2; i++ {
            for j := i + 1; j < L - 1; j++ {
                for k := j + 1; k < L; k++ {
                    if nums[h] + nums[i] + nums[j] + nums[k] == target {
                        tmp := []int{nums[h], nums[i], nums[j], nums[k]}
                        ans = append(ans, tmp)
                    }
                }
            }
        }
    }
    // fmt.Println(ans)

    return ans
}
