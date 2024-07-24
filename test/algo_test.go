package test

import (
	"Push-Swap/algorithms"
	"Push-Swap/stack"
	"fmt"
	"testing"
)

func TestAlgo(t *testing.T) {
	s := stack.Stack{
		StackA: []int{7, 6, 2, 1, 3},
	}
	excpected := []int{1, 2, 3, 6, 7}

	fmt.Println("Before:", s.StackA)
	if !algorithms.SortAlgorithm(&s) {
		t.Fatalf("This stack could not be sorted.")
	}
	fmt.Println("After:", s.StackA)
	fmt.Printf("%d Instructions: %v\n", s.InstructionCount, s.Operation)

	for i, v := range s.StackA {
		if v != excpected[i] {
			t.Fatalf("Have: %v; Want: %v", s.StackA, excpected)
		}
	}

	s = stack.Stack{
		StackA: []int{5, 49, 21, 75, 32, 64, 98},
	}
	excpected = []int{5, 21, 32, 49, 64, 75, 98}

	fmt.Println("Before:", s.StackA)
	if !algorithms.SortAlgorithm(&s) {
		t.Fatalf("This stack could not be sorted.")
	}
	fmt.Println("After:", s.StackA)
	fmt.Printf("%d Instructions: %v\n", s.InstructionCount, s.Operation)

	for i, v := range s.StackA {
		if v != excpected[i] {
			t.Fatalf("Have: %v; Want: %v", s.StackA, excpected)
		}
	}

	s = stack.Stack{
		StackA: []int{71, 48, 29, 79, 3, 26, 13, 9, 52, 68, 59, 92, 84},
	}
	excpected = []int{3, 9, 13, 26, 29, 48, 52, 59, 68, 71, 79, 84, 92}

	fmt.Println("Before:", s.StackA)
	if !algorithms.SortAlgorithm(&s) {
		t.Fatalf("This stack could not be sorted.")
	}
	fmt.Println("After:", s.StackA)
	fmt.Printf("%d Instructions: %v\n", s.InstructionCount, s.Operation)

	for i, v := range s.StackA {
		if v != excpected[i] {
			t.Fatalf("Have: %v; Want: %v", s.StackA, excpected)
		}
	}
}
