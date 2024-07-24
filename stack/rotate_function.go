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
}

func (s *Stack) Rb() {
	if len(s.StackB) < 2 {
		fmt.Println("Error: not enough arguments to rotate B")
	}
	first := s.StackB[0]

	s.StackB = s.StackB[1:]

	s.StackB = append(s.StackB, first)
}

func (s *Stack) Rr() {
	if len(s.StackA) >= 2 {
		first := s.StackA[0]
		s.StackA = s.StackA[1:]
		s.StackA = append(s.StackA, first)
	}

	if len(s.StackB) >= 2 {
		first := s.StackB[0]
		s.StackB = s.StackB[1:]
		s.StackB = append(s.StackB, first)
	}
}
