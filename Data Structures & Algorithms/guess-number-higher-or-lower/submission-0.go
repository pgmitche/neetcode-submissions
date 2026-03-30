/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    lhs, rhs := 1, n

	for lhs <= rhs {
		guessNum := (lhs + rhs) / 2

		switch result := guess(guessNum); result{
			case -1:
				rhs = guessNum - 1
			case 1:
				lhs = guessNum + 1
			case 0:
				return guessNum
		}
	}

	return -1
}
