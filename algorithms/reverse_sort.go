package algorithm

import (
	instructions "Push-Swap/instruction"
	"Push-Swap/stack"
)

func ReverseSort(s *stack.Stack) {
	maxA := Max(s.StackA)
	middle := len(s.StackA) / 2
	for i := 0; i < middle; i++ {
		instructions.ExecuteInstruction(s, "rra")
	}

	for i := 0; i < middle; i++ {
		instructions.ExecuteInstruction(s, "pb")
	}
	maxB := Max(s.StackB)

	if maxA == s.StackA[0] && maxB != s.StackB[0] {
		instructions.ExecuteInstruction(s, "rr")
	}
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
