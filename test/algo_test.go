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
	algorithms.SortAlgorithm(&s)
	fmt.Println("After:", s.StackA)

	for i, v := range s.StackA {
		if v != excpected[i] {
			t.Fatalf("Have: %v; Want: %v", s.StackA, excpected)
		}
	}
}
