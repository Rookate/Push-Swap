package stack

import "fmt"

func (s *Stack) Rra() {
	if len(s.StackA) < 2 {
		fmt.Println("Error: not enough arguments to reverse rotate A")
		return
	}

	last := s.StackA[len(s.StackA)-1]
	s.StackA = s.StackA[:len(s.StackA)-1]
	s.StackA = append([]int{last}, s.StackA...)
}

func (s *Stack) Rrb() {
	if len(s.StackB) < 2 {
		fmt.Println("Error: not enough arguments to revers rotate B")
		return
	}

	last := s.StackB[len(s.StackB)-1]
	s.StackB = s.StackB[:len(s.StackB)-1]
	s.StackB = append([]int{last}, s.StackB...)
}

func (s *Stack) Rrr() {
	s.Rra()
	s.Rrb()
}
