package instructions

import (
	"Push-Swap/stack"
	"testing"
)

func TestExecuteInstruction(t *testing.T) {
	t.Run("Test sa", func(t *testing.T) {
		s := stack.Stack{
			StackA: []int{2, 1, 3},
			StackB: []int{},
		}
		ExecuteInstruction(&s, "sa")

		expected := []int{2, 3, 1}
		for i, v := range s.StackA {
			if v != expected[i] {
				t.Errorf("Expected StackA to be %v, got %v", expected, s.StackA)
			}
		}
	})

	t.Run("Test pb", func(t *testing.T) {
		s := stack.Stack{
			StackA: []int{1, 2, 3},
			StackB: []int{},
		}
		ExecuteInstruction(&s, "pb")

		expectedA := []int{1, 2}
		expectedB := []int{3}
		for i, v := range s.StackA {
			if v != expectedA[i] {
				t.Errorf("Expected StackA to be %v, got %v", expectedA, s.StackA)
			}
		}
		for i, v := range s.StackB {
			if v != expectedB[i] {
				t.Errorf("Expected StackB to be %v, got %v", expectedB, s.StackB)
			}
		}
	})

	t.Run("Test ra", func(t *testing.T) {
		s := stack.Stack{
			StackA: []int{1, 2, 3},
			StackB: []int{},
		}
		ExecuteInstruction(&s, "ra")

		expected := []int{2, 3, 1}
		for i, v := range s.StackA {
			if v != expected[i] {
				t.Errorf("Expected StackA to be %v, got %v", expected, s.StackA)
			}
		}
	})

	t.Run("Test rra", func(t *testing.T) {
		s := stack.Stack{
			StackA: []int{1, 2, 3},
			StackB: []int{},
		}
		ExecuteInstruction(&s, "rra")

		expected := []int{3, 1, 2}
		for i, v := range s.StackA {
			if v != expected[i] {
				t.Errorf("Expected StackA to be %v, got %v", expected, s.StackA)
			}
		}
	})
}
