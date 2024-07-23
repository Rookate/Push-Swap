package algorithms

import (
	"Push-Swap/instructions"
	"Push-Swap/stack"
)

func SortThree(s *stack.Stack) {
	if len(s.StackA) != 3 {
		return
	}

	a := s.StackA[0]
	b := s.StackA[1]
	c := s.StackA[2]

	if a > b && b > c {
		instructions.ExecuteInstruction(s, "sa")
		instructions.ExecuteInstruction(s, "rra")
	} else if a > c && c > b {
		instructions.ExecuteInstruction(s, "ra")
	} else if b > a && a > c {
		instructions.ExecuteInstruction(s, "sa")
	} else if b > c && c > a {
		instructions.ExecuteInstruction(s, "rra")
	} else if c > a && a > b {
		instructions.ExecuteInstruction(s, "sa")
		instructions.ExecuteInstruction(s, "ra")
	}
}
