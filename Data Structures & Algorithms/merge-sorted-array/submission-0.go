func merge(nums1 []int, m int, nums2 []int, n int) {
	tempNums1 := make([]int, m)
    copy(tempNums1, nums1[:m])

	lhsCursor := 0
	rhsCursor := 0

	// Move the rhs cursor along as we overwrite nums1
	for cursor := 0; cursor < m+n; cursor++ {
		// so long as the lhsCursor is inbounds AND
		// if the lhsCursor value is less than or equal to
		// the value at the rhsCursor, insert and move.
		// OR if rhsCursor has moved out of bounds
	    if lhsCursor < m && (rhsCursor >= n || tempNums1[lhsCursor] <= nums2[rhsCursor]) {
			nums1[cursor] = tempNums1[lhsCursor]
			lhsCursor++
		} else {
			// if the rhs value is smaller, then insert
			// that and shift the rhs cursor
			nums1[cursor] = nums2[rhsCursor]
			rhsCursor++
		}
	}
}
