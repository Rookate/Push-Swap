package stack

import (
	"fmt"
	"testing"
)

func TestRa(t *testing.T) {
	s := Stack{
		StackA: []int{1, 2, 3, 4, 5},
	}

	// Affiche la pile avant l'exécution de la méthode Ra
	fmt.Println("Avant Ra:", s.StackA)

	// Exécute la méthode Ra
	s.Ra()

	// Affiche la pile après l'exécution de la méthode Ra
	fmt.Println("Après Ra:", s.StackA)

	// Valeur attendue
	expected := []int{2, 3, 4, 5, 1}
	fmt.Println("Expected:", expected)

	// Vérifie si le résultat est correct et affiche une erreur si ce n'est pas le cas
	for i, v := range s.StackA {
		if v != expected[i] {
			t.Errorf("Ra() = %v; want %v", s.StackA, expected)
		}
	}
}
func TestRb(t *testing.T) {
	s := Stack{
		StackB: []int{5, 4, 3, 2, 1},
	}

	// Affiche la pile avant l'exécution de la méthode Rb
	fmt.Println("Avant Rb:", s.StackB)

	// Exécute la méthode Rb
	s.Rb()

	// Affiche la pile après l'exécution de la méthode Rb
	fmt.Println("Après Rb:", s.StackB)

	// Valeur attendue
	expected := []int{4, 3, 2, 1, 5}
	fmt.Println("Expected:", expected)

	// Vérifie si le résultat est correct et affiche une erreur si ce n'est pas le cas
	for i, v := range s.StackB {
		if v != expected[i] {
			t.Errorf("Rb() = %v; want %v", s.StackB, expected)
		}
	}
}

func TestRR(t *testing.T) {
	s := Stack{
		StackA: []int{1, 2, 5, 0, 7},
		StackB: []int{5, 4, 3, 2, 1},
	}

	expectedA := []int{2, 5, 0, 7, 1}
	expectedB := []int{4, 3, 2, 1, 5}

	fmt.Println("Avant Ra:", s.StackA)
	fmt.Println("Avant Rb", s.StackB)

	s.Rr()

	fmt.Println("Après Ra:", s.StackA)
	fmt.Println("Expected Ra:", expectedA)
	fmt.Println("Après Rb:", s.StackB)
	fmt.Println("Ecpected Rb:", expectedB)

	for i, v := range s.StackA {
		if v != expectedA[i] {
			t.Errorf("Ra() = %v; want %v", s.StackA, expectedA)
		}
	}

	for i, v := range s.StackB {
		if v != expectedB[i] {
			t.Errorf("Rb() = %v; want : %v", s.StackB, expectedB)
		}
	}
}

func TestRra(t *testing.T) {
	s := Stack{
		StackA: []int{1, 2, 3, 4, 5},
	}

	fmt.Println("Avant Rra:", s.StackA)
	s.Rra()
	fmt.Println("Après Rra:", s.StackA)

	expected := []int{5, 1, 2, 3, 4}
	fmt.Println("Expected:", expected)

	for i, v := range s.StackA {
		if v != expected[i] {
			t.Errorf("Rra() = %v; want %v", s.StackA, expected)
		}
	}
}

func TestRrb(t *testing.T) {
	s := Stack{
		StackB: []int{1, 2, 7, 4, 5},
	}

	fmt.Println("Avant Rra:", s.StackB)
	s.Rrb()
	fmt.Println("Après Rra:", s.StackB)

	expected := []int{5, 1, 2, 7, 4}
	fmt.Println("Expected:", expected)

	for i, v := range s.StackB {
		if v != expected[i] {
			t.Errorf("Rra() = %v; want %v", s.StackB, expected)
		}
	}
}
