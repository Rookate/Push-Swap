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
	s.Operation = append(s.Operation, instruction)
	s.InstructionCount++
	return nil
}

func RollbackInstruction(s *stack.Stack) error {
	switch s.Operation[len(s.Operation)-1] {
	case "sa":
		s.Sa()
	case "sb":
		s.Sb()
	case "ss":
		s.Ss()
	case "pa":
		s.Pb()
	case "pb":
		s.Pa()
	case "ra":
		s.Rra()
	case "rb":
		s.Rrb()
	case "rr":
		s.Rrr()
	case "rra":
		s.Ra()
	case "rrb":
		s.Rb()
	case "rrr":
		s.Rr()
	default:
		return errors.New("invalid instruction")
	}
	s.Operation = s.Operation[:len(s.Operation)-1]
	s.InstructionCount--
	return nil
}
