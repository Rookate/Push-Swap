package stack

import (
	"fmt"
)

type Stack struct {
	StackA           []int
	StackB           []int
	InstructionCount int
	Operation        []string
}

func (s *Stack) PrintStack() {
	fmt.Println("Stack A:", s.StackA)
	fmt.Println("Stack B:", s.StackB)
	fmt.Println("Instruction Count:", s.InstructionCount)
	fmt.Println("Flags:", s.Operation)
}
