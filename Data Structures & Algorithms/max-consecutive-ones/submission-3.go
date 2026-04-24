func findMaxConsecutiveOnes(nums []int) int {
    maxCounter := 0
    counter := 0
    for _, n := range nums {
        if n == 1 {
            counter++ 
            if counter > maxCounter {
                maxCounter = counter
            }
        } else { 
            counter = 0
        } 
    }
    return maxCounter 
}
