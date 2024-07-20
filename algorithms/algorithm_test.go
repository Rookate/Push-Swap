package algorithm

import (
	"Push-Swap/stack"
	"fmt"
	"testing"
)

func TestReverseSort(t *testing.T) {
	s := stack.Stack{
		StackA: []int{122, 80, 36, 20},
		StackB: []int{},
	}

	fmt.Println("Before func:", s.StackA)
	fmt.Println("Before func:", s.StackB)
	Reseolve(&s)
	fmt.Println("After func:", s.StackA)
	fmt.Println("After func:", s.StackB)
	fmt.Println("Number of instructions:", s.InstructionCount)
	fmt.Println("Operations:", s.Operation)

	expectedA := []int{20, 36, 80, 122}
	expectedB := []int{}

	if len(s.StackA) != len(expectedA) {
		t.Errorf("Have: %v; want: %v", s.StackA, expectedA)
		return
	}

	for i, v := range s.StackA {
		if v != expectedA[i] {
			t.Errorf("Have: %v; want: %v", s.StackA, expectedA)
		}
	}

	if len(s.StackB) != len(expectedB) {
		t.Errorf("Have: %v; want: %v", s.StackB, expectedB)
		return
	}

	for i, v := range s.StackB {
		if v != expectedB[i] {
			t.Errorf("Have: %v; want: %v", s.StackB, expectedB)
		}
	}
}
