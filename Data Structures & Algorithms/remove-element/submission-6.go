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
    j := len(nums) - 1
    
    for i <= j {
        // ensure that the end isn't a val
        for nums[j] == val && j > 0 {
            j--
        }
        if nums[i] == val {
            nums[i] = nums[j]
            j--
        }
        i++
    }

    // duck-tape this solution so it works 
    // basically if the first element is a val, pretend like we didn't make an oopsie
    if i > 0 && nums[0] == val {
        return 0
    }

    return i 
}
