package instructions

import (
	"Push-Swap/stack"
	"errors"
)

// ExecuteInstruction applique une instruction sur la pile
func ExecuteInstruction(s *stack.Stack, instruction string) error {
	switch instruction {
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
		return errors.New("invalid instruction")
	}
	s.InstructionCount++
	return nil
}
