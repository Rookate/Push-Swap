package stack

import (
	"fmt"
)

type Stack struct {
	StackA []int
	StackB []int
}

func (s *Stack) PrintStack() {
	fmt.Println("Stack A:", s.StackA)
	fmt.Println("Stack B:", s.StackB)
}
