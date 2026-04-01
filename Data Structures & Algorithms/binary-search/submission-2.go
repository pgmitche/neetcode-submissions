func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// given that the nums slice is sorted, split the problem
	// space in half by finding the midpoint and checking for 
	// equivalence, then providing the closer half of the slice
	// to recursive nums
	var mid int 
	mid = len(nums) / 2

	if nums[mid] > target {
		return search(nums[:mid], target)
	} else if nums[mid] < target {
		// to keep the recursive subslicing approachmid will be incorrect 
		// if we don't combine the offset for the RHS approach
		// using lhs, rhs cursors or a closure is cleaner than this
		result := search(nums[mid+1:], target)
		if result == -1 {
			return -1
		}
		return result + mid + 1
	} else{
		return mid
	}
}
