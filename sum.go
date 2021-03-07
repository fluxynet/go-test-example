package gote

// Sum of numbers
func Sum(positives bool, nums ...int) int {
	var t int
	for i := range nums {
		if positives && nums[i] <= 0 {
			continue
		}

		t += nums[i]
	}

	return t
}
