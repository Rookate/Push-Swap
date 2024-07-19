package stack

import "fmt"

func (s *Stack) Pa() {
	if len(s.StackA) == 0 {
		fmt.Println("Error: Pile A is empty")
		return
	}

	value := s.StackA[len(s.StackA)-1]
	s.StackA = s.StackA[:len(s.StackA)-1]

	s.StackB = append(s.StackB, value)
	s.Operation = append(s.Operation, "pa")
	s.InstructionCount++
}

func (s *Stack) Pb() {
	if len(s.StackB) == 0 {
		fmt.Println("Error : Pile B is empty")
		return
	}

	value := s.StackB[len(s.StackB)-1]
	s.StackB = s.StackB[:len(s.StackB)-1]

	s.StackA = append(s.StackA, value)

	s.Operation = append(s.Operation, "pb")
	s.InstructionCount++
}
