func search(nums []int, target int) int {
	lhs, rhs := 0, len(nums)-1

	// while the cursors are apart or equivalent
	for lhs <= rhs {
		
		// identify the seeker
		seeker := (lhs + rhs) / 2

		// is the target lower or higher than the seeker
		// always move your seeker toward the target!
		if nums[seeker] > target{
			rhs = seeker-1
		} else if nums[seeker] < target{
			lhs = seeker+1
		} else {
			// seeker aligned on target
			return seeker
		}
	}

	// target not found
	return -1
}
