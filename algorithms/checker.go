package algorithms

import (
	"Push-Swap/instructions"
	"Push-Swap/stack"
)

// Checker vérifie si la pile est triée.
func Checker(s *stack.Stack) string {
	if IsSorted(s.StackA) {
		return "OK"
	}
	return "KO"
}

func ExecuteChecker(s *stack.Stack, instruction string) {
	instructions.ExecuteInstruction(s, instruction)
}
