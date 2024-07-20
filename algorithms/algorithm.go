package algorithm

import (
	"Push-Swap/stack"
	"fmt"
)

func Reseolve(s *stack.Stack) {
	if len(s.StackA) == 0 {
		fmt.Println("Error: Stack A is empty")
		return
	}

	if IsSorted(s.StackA) {
		fmt.Println("Stack already sort")
		return
	}

	if IsDescending(s.StackA) {
		ReverseSort(s)
	}
}
