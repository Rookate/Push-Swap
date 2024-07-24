package stack

import "fmt"

func (s *Stack) Sa() {
	if len(s.StackA) < 2 {
		fmt.Println("Error : not enough element to swap")
	}

	first := s.StackA[0]
	second := s.StackA[1]

	s.StackA[0] = second
	s.StackA[1] = first
}

func (s *Stack) Sb() {
	if len(s.StackB) < 2 {
		fmt.Println("Error: not enough elements to swap")
		return
	}

	first := s.StackB[0]
	second := s.StackB[1]

	s.StackB[0] = second
	s.StackB[1] = first
}

func (s *Stack) Ss() {
	if len(s.StackA) >= 2 {
		first := s.StackA[0]
		second := s.StackA[1]

		s.StackA[0] = second
		s.StackA[1] = first
	}

	if len(s.StackB) >= 2 {
		first := s.StackB[0]
		second := s.StackB[1]

		s.StackB[0] = second
		s.StackB[1] = first
	}
}
