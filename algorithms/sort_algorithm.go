package algorithms

import (
	"Push-Swap/instructions"
	"Push-Swap/stack"
)

func SortAlgorithm(s *stack.Stack) bool {
	if IsSorted(s.StackA) && len(s.StackB) == 0 {
		return true
	}

	var possibles []string

	if IsSorted(s.StackA) && IsDescending(s.StackB) {
		possibles = append(possibles, "pa")
	}

	// What can we do here?
	if len(s.StackA) > 1 && len(s.StackB) > 1 {
		if len(s.StackA) > 1 && len(s.StackB) > 2 {
			if s.StackA[len(s.StackA)-1] == Min(s.StackA) && s.StackB[len(s.StackB)-1] == Max(s.StackB) {
				possibles = append(possibles, "rrr")
			}
			if s.StackA[0] == Max(s.StackA) && s.StackB[0] == Min(s.StackB) {
				possibles = append(possibles, "rr")
			}
		}
		if s.StackA[0] > s.StackA[1] && s.StackB[0] < s.StackB[1] {
			possibles = append(possibles, "ss")
		}
	}
	if len(s.StackA) > 1 {
		if len(s.StackA) > 2 {
			if s.StackA[len(s.StackA)-1] == Min(s.StackA) {
				possibles = append(possibles, "rra")
			}
			if s.StackA[0] == Max(s.StackA) {
				possibles = append(possibles, "ra")
			}
		}
		if s.StackA[0] > s.StackA[1] {
			possibles = append(possibles, "sa")
		}
	}

	if len(s.StackB) > 0 {
		if len(s.StackB) > 1 {
			if len(s.StackB) > 2 {
				if s.StackB[len(s.StackB)-1] == Max(s.StackB) {
					possibles = append(possibles, "rrb")
				}
				if s.StackB[0] == Min(s.StackB) {
					possibles = append(possibles, "rb")
				}
			}
			if s.StackB[0] < s.StackB[1] {
				possibles = append(possibles, "sb")
			}
		}
		if IsSorted(s.StackA) && s.StackB[0] < s.StackA[0] {
			possibles = append(possibles, "pa")
		}
	}

	if len(s.StackA) > 0 {
		possibles = append(possibles, "pb")
	}

	// Let's do each and find out...
	for _, inst := range possibles {
		instructions.ExecuteInstruction(s, inst)
		// TODO : how do I call the function itself?
		done := SortAlgorithm(s)
		if !done {
			instructions.RollbackInstruction(s)
		} else {
			return done
		}
	}

	return false
}
