func replaceElements(arr []int) []int {
	max := -1
	for i := len(arr)-1; i >= 0; i-- {
		val := arr[i]
		arr[i] = max
		if val > max {
			max = val
		}
	}

	return arr
}
