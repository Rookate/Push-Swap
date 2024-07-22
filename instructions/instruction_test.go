package instructions

import (
	"Push-Swap/stack"
	"fmt"
	"testing"
)

func TestExecuteInstruction(t *testing.T) {
	t.Run("Test sa", func(t *testing.T) {
		s := stack.Stack{
			StackA: []int{2, 1, 3},
			StackB: []int{},
		}
		ExecuteInstruction(&s, "sa")

		expected := []int{1, 2, 3}
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

		expectedA := []int{2, 3}
		expectedB := []int{1}
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

func TestAll(t *testing.T) {
	s := stack.Stack{
		StackA: []int{2, 1, 3, 6, 5, 8},
	}

	fmt.Println("Avant Prog:", s.StackA)
	s.Sa()
	s.Pb()
	s.Pb()
	s.Pb()
	s.Sa()
	s.Pa()
	s.Pa()
	s.Pa()
	fmt.Println("Après Prog:", s.StackA)
	fmt.Println("Flags:", s.InstructionCount)

	expected := []int{1, 2, 3, 5, 6, 8}
	fmt.Println("Expected:", expected)

	for i, v := range s.StackA {
		if v != expected[i] {
			t.Errorf("Prog() = %v; want %v", s.StackA, expected)
		}
	}
}

func TestAllsol2(t *testing.T) {
	s := stack.Stack{
		StackA: []int{7, 1, 3, 2, 4},
		StackB: []int{},
	}

	fmt.Println("Avant Prog stack a:", s.StackA)
	fmt.Println("Avant Prog stack b:", s.StackB)
	s.Pb()
	s.Pb()
	s.Ss()
	s.Pa()
	s.Ra()
	s.Pa()
	fmt.Println("Après Prog stack a:", s.StackA)
	fmt.Println("Après Prog stack b:", s.StackB)
	fmt.Println("Number of Instruction:", s.InstructionCount)
	fmt.Println("Flags:", s.Operation)

	expected := []int{1, 2, 3, 4, 7}
	fmt.Println("Expected:", expected)

	for i, v := range s.StackA {
		if v != expected[i] {
			t.Errorf("Prog() = %v; want %v", s.StackA, expected)
		}
	}
}

func TestAllsol3(t *testing.T) {
	s := stack.Stack{
		StackA: []int{4, 2, 1, 5, 3},
		StackB: []int{},
	}

	fmt.Println("Avant Prog stack a:", s.StackA)
	fmt.Println("Avant Prog stack b:", s.StackB)
	s.Pb()
	s.Pb()
	s.Pb()
	s.Rr()
	s.Sb()
	s.Pa()
	s.Sa()
	s.Pa()
	s.Pa()
	fmt.Println("Après Prog stack a:", s.StackA)
	fmt.Println("Après Prog stack b:", s.StackB)
	fmt.Println("Number of Instruction:", s.InstructionCount)
	fmt.Println("Flags:", s.Operation)

	expected := []int{1, 2, 3, 4, 5}
	fmt.Println("Expected:", expected)

	for i, v := range s.StackA {
		if v != expected[i] {
			t.Errorf("Prog() = %v; want %v", s.StackA, expected)
		}
	}
}
