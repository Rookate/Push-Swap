package algorithms

func IsSorted(stack []int) bool {
	if len(stack) <= 1 {
		return true
	}

	for i := 1; i < len(stack); i++ {
		if stack[i] < stack[i-1] {
			return false
		}
	}
	return true
}

func IsDescending(stack []int) bool {
	for i := 0; i < len(stack)-1; i++ {
		if stack[i] < stack[i+1] {
			return false
		}
	}
	return true
}
func Max(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	max := slice[0]

	for _, value := range slice[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

func Min(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	min := slice[0]

	for _, value := range slice[1:] {
		if value < min {
			min = value
		}
	}
	return min
}
