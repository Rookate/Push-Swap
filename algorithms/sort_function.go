package algorithm

func IsSorted(stack []int) bool {
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
