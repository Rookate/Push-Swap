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
}

func IsSorted(stack []int) bool {
	for i := 1; i < len(stack); i++ {
		if stack[i] < stack[i-1] {
			return false
		}
	}
	return true
}
