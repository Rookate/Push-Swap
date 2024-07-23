package algorithms

import (
	"Push-Swap/instructions"
	"Push-Swap/stack"
)

func SortAlgorithm(s *stack.Stack) {
	middle := len(s.StackA) / 2
	if len(s.StackA)%2 != 0 {
		middle++
	}

	for !IsSorted(s.StackA) || len(s.StackB) != 0 {
		if len(s.StackA) > 1 && s.StackA[0] == Max(s.StackA) {
			instructions.ExecuteInstruction(s, "ra")
			continue
		} else if len(s.StackA) > 1 && Min(s.StackA) == s.StackA[len(s.StackA)-1] {
			instructions.ExecuteInstruction(s, "rra")
			continue
		} else if len(s.StackA) > 1 && s.StackA[0] > s.StackA[1] {
			instructions.ExecuteInstruction(s, "sa")
			continue
		} else {
			if len(s.StackA) > 1 {
				for i := 0; i < middle; i++ {
					instructions.ExecuteInstruction(s, "pb")
				}
			}

			for len(s.StackB) > 0 {
				if len(s.StackB) == 1 && IsSorted(s.StackA) {
					instructions.ExecuteInstruction(s, "pa")
					break
				}
				if len(s.StackB) > 1 || !IsSorted(s.StackA) {
					if len(s.StackA) > 1 && len(s.StackB) > 1 && s.StackA[0] > s.StackA[1] && s.StackA[0] != Max(s.StackA) && s.StackB[0] < s.StackB[1] {
						instructions.ExecuteInstruction(s, "ss")
						continue
					} else if len(s.StackA) > 1 && s.StackA[0] > s.StackA[1] && s.StackA[0] != Max(s.StackA) {
						instructions.ExecuteInstruction(s, "sa")
						continue
					} else if len(s.StackB) > 1 && s.StackB[0] < s.StackB[1] {
						instructions.ExecuteInstruction(s, "sb")
						continue
					} else if len(s.StackA) > 1 && len(s.StackB) > 1 && s.StackA[0] == Max(s.StackA) && s.StackB[0] == Min(s.StackB) {
						instructions.ExecuteInstruction(s, "rr")
						continue
					} else if len(s.StackA) > 1 && s.StackA[0] == Max(s.StackA) {
						instructions.ExecuteInstruction(s, "ra")
						continue
					} else if len(s.StackB) > 1 && s.StackB[0] == Min(s.StackB) {
						instructions.ExecuteInstruction(s, "rb")
						continue
					} else if len(s.StackA) > 1 && len(s.StackB) > 1 && s.StackA[len(s.StackA)-1] == Min(s.StackA) && s.StackB[len(s.StackB)-1] == Max(s.StackB) {
						instructions.ExecuteInstruction(s, "rrr")
						continue
					} else if len(s.StackA) > 1 && s.StackA[len(s.StackA)-1] == Min(s.StackA) {
						instructions.ExecuteInstruction(s, "rra")
						continue
					} else if len(s.StackB) > 1 && s.StackB[len(s.StackB)-1] == Max(s.StackB) {
						instructions.ExecuteInstruction(s, "rrb")
						continue
					} else {
						instructions.ExecuteInstruction(s, "pa")
						continue
					}
				}
			}
		}
	}
}
