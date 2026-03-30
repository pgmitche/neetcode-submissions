func searchMatrix(matrix [][]int, target int) bool {
	// align on rows
	rows, cols := len(matrix),  len(matrix[0])
	lhs, rhs := 0, rows*cols - 1

	// by treating the whole 2d matrix as a single array, it can be indexed by
	// dividing the seeker index by the matrix shape
	// and fetching the remainder from the matrix shape
	// 1D position = matrix[seeker/length][seeker%length]

	// while the cursors are apart or equivalent
    for lhs <= rhs {
        // identify the seeker
        seeker := (lhs + rhs) / 2
        // is the target lower or higher than the seeker
        // always move your seeker toward the target!
		val := matrix[seeker/cols][seeker%cols] 
        if val > target{
            rhs = seeker-1
        } else if val < target{
            lhs = seeker+1
        } else {
            // seeker aligned on target
            return true
        }
    }

    return false
}
