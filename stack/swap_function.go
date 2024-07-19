package stack

import "fmt"

func (s *Stack) Sa() {
	if len(s.StackA) < 2 {
		fmt.Println("Error : not enough element to swap")
	}

	first := s.StackA[len(s.StackA)-1]
	second := s.StackA[len(s.StackA)-2]

	s.StackA = s.StackA[:len(s.StackA)-2]

	s.StackA = append(s.StackA, first, second)
	s.Operation = append(s.Operation, "sa")
	s.InstructionCount++

}

func (s *Stack) Sb() {
	if len(s.StackB) < 2 {
		fmt.Println("Error: not enough elements to swap")
		return
	}

	first := s.StackB[len(s.StackB)-1]
	second := s.StackB[len(s.StackB)-2]

	s.StackB = s.StackB[:len(s.StackB)-2]

	s.StackB = append(s.StackB, first, second)
	s.Operation = append(s.Operation, "sb")
	s.InstructionCount++
}

func (s *Stack) Ss() {
	s.Sa()
	s.Sb()
	s.Operation = append(s.Operation, "ss")
	s.InstructionCount++
}
