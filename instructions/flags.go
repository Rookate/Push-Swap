package instructions

import (
	"Push-Swap/stack"
	"fmt"
)

func ExecuteInstruction(s *stack.Stack, instructions string) {
	switch instructions {
	case "sa":
		s.Sa()
	case "sb":
		s.Sb()
	case "ss":
		s.Ss()
	case "pa":
		s.Pa()
	case "pb":
		s.Pb()
	case "ra":
		s.Ra()
	case "rb":
		s.Rb()
	case "rr":
		s.Rr()
	case "rra":
		s.Rra()
	case "rrb":
		s.Rrb()
	case "rrr":
		s.Rrr()
	default:
		fmt.Println("Error: unknown instruction", instructions)
		return
	}
	s.InstructionCount++
}
