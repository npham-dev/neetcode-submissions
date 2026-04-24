import (
    "slices"
)

func removeElement(nums []int, val int) int {
    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] == val {
            nums = slices.Delete(nums, i, i + 1)
        }
    }
    return len(nums)
}
