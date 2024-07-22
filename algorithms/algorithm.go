package algorithms

import (
	"Push-Swap/stack"
	"fmt"
)

func Resolve(s *stack.Stack) {
	if len(s.StackA) == 0 {
		fmt.Println("Error: Stack A is empty")
		return
	}

	if IsSorted(s.StackA) {
		fmt.Println("Stack already sort")
		return
	}
	if IsDescending(s.StackA) && len(s.StackA) <= 5 {
		ReverseSort(s)
	}
	SortAlgorithm(s)
}
