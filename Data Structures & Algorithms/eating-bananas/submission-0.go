func minEatingSpeed(piles []int, h int) int {
	// get max in piles
	max := 0 
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}

	lhs, rhs := 1, max 
	minSpeed := rhs
	for lhs <= rhs {
		guessRate := (lhs + rhs) / 2 

		var checkHours int
		for _, pile := range piles{
			// Specify cieling division for Go, default is floor div
			checkHours += (pile + guessRate - 1) / guessRate
		}

		if checkHours <= h {
			if guessRate < minSpeed {
				minSpeed = guessRate
			}

			rhs = guessRate - 1
		} else {
			lhs = guessRate + 1
		}
	}

	return minSpeed
}

