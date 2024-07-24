package stack

import "fmt"

func (s *Stack) Pb() {
	if len(s.StackA) == 0 {
		fmt.Println("Error: Pile A is empty")
		return
	}

	value := s.StackA[0]
	s.StackA = s.StackA[1:]

	s.StackB = append([]int{value}, s.StackB...)
}

func (s *Stack) Pa() {
	if len(s.StackB) == 0 {
		fmt.Println("Error : Pile B is empty")
		return
	}

	value := s.StackB[0]
	s.StackB = s.StackB[1:]

	s.StackA = append([]int{value}, s.StackA...)
}
