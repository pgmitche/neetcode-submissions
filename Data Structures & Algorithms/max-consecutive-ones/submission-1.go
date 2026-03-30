func findMaxConsecutiveOnes(nums []int) int {
	var max int = 0
    counter := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 1 {
            counter = 0
            continue
        }
        counter++
        if counter > max {
            max = counter
        }
    }
    return max
}
