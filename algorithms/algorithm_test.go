package algorithms

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
	Resolve(&s)
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

func TestMin(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{3, 1, 4, 1, 5, 9, 2, 6, 5}, 1},
		{[]int{10, 20, 30}, 10},
		{[]int{7, 7, 7, 7, 7}, 7},
		{[]int{-5, -10, 0, 5, 10}, -10},
		{[]int{}, 0}, // Testing the empty slice case
	}
	for _, test := range tests {
		result := Min(test.input)
		if result != test.expected {
			t.Errorf("Min(%v) = %d; expected %d", test.input, result, test.expected)
		} else {
			fmt.Println("Result:", result)
			fmt.Println("Expected:", test.expected)
		}
	}
}
