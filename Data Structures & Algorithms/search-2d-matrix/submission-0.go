func searchMatrix(matrix [][]int, target int) bool {
	// align on rows
	lhs, rhs := 0, len(matrix[0]) - 1

	rowIndex := 0
	for  {
		// if no row contains the value, fail
		if rowIndex > len(matrix)-1 {
			return false
		}

		// if the value is within bounds then break
		if matrix[rowIndex][lhs] <= target && target <= matrix[rowIndex][rhs] {
			break
		}

		rowIndex++
	}

	// then align on columns
	targetRow := matrix[rowIndex]

	// while the cursors are apart or equivalent
	for lhs <= rhs {
		
		// identify the seeker
		seeker := (lhs + rhs) / 2

		// is the target lower or higher than the seeker
		// always move your seeker toward the target!
		if targetRow[seeker] > target{
			rhs = seeker-1
		} else if targetRow[seeker] < target{
			lhs = seeker+1
		} else {
			// seeker aligned on target
			return true
		}
	}

	return false
}
