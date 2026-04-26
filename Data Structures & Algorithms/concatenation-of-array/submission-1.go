func getConcatenation(nums []int) []int {
    length := len(nums)
    ans := make([]int, length * 2)
    for i := 0; i < length; i++ {
        ans[i] = nums[i]
        ans[length + i] = nums[i]
    } 
    return ans
}
