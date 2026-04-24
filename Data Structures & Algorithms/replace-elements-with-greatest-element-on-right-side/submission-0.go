func replaceElements(arr []int) []int {
	rightMax := -1	
	for i := len(arr) - 1; i >= 0; i-- {
		temp := arr[i]
		arr[i] = rightMax
		if temp > rightMax {
			rightMax = temp
		}
	}
	return arr
}

