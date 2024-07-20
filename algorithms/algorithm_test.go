package algorithm

import (
	"Push-Swap/stack"
	"fmt"
	"testing"
)

func TestReverseSort(t *testing.T) {
	s := stack.Stack{
		StackA: []int{5, 4, 3, 2, 1},
		StackB: []int{},
	}

	fmt.Println("Before func:", s.StackA)
	fmt.Println("Before func:", s.StackB)
	ReverseSort(&s)
	fmt.Println("After func:", s.StackA)
	fmt.Println("After func:", s.StackB)
	fmt.Println("Number of instructions:", s.InstructionCount)
	fmt.Println("Operations:", s.Operation)

	expectedA := []int{4, 3, 5}
	expectedB := []int{2, 1}

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
