func removeElement(nums []int, val int) int {
    // for i := len(nums) - 1; i >= 0; i-- {
    //     if nums[i] == val {
    //         nums = slices.Delete(nums, i, i + 1)
    //     }
    // }
    // return len(nums)

    // the above works, but there's a better way that doesn't involve shifting
    // because we only care about the first k = non-val elements, we can just use swaps instead

    i := 0
    j := len(nums)
    
    for i < j {
        if nums[i] == val {
            j--
            nums[i] = nums[j]
        } else {
            i++
        }
    }

    return j 
}
