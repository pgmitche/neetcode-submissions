func sortColors(nums []int) {
    bucket := make([]int, 3)

	// create a bucket array using the colour code matching the index
	for _, colour := range nums{
		bucket[colour]++
	}
	
	// for each colour in the bucket
	cursor := 0
	for colour, colourCount := range bucket {
		// for each count of the colour in the bucket
		for count := 0; count < colourCount; count++ {
			// insert the colour code at the colour index with an offset of the count
			nums[cursor] = colour
			cursor++
		}
	}
}
