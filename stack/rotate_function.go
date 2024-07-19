package stack

import "fmt"

func (s *Stack) Ra() {
	if len(s.StackA) < 2 {
		fmt.Println("Error: not enough arguments to rotate A")
		return
	}

	first := s.StackA[0]

	s.StackA = s.StackA[1:]

	s.StackA = append(s.StackA, first)
	s.Operation = append(s.Operation, "ra")
	s.InstructionCount++
}

func (s *Stack) Rb() {
	if len(s.StackB) < 2 {
		fmt.Println("Error: not enough arguments to roate B")
	}
	first := s.StackB[0]

	s.StackB = s.StackB[1:]

	s.StackB = append(s.StackB, first)
	s.Operation = append(s.Operation, "rb")
	s.InstructionCount++
}

func (s *Stack) Rr() {
	s.Ra()
	s.Rb()
	s.Operation = append(s.Operation, "rr")
	s.InstructionCount++
}
