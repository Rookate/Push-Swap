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

func TestSa(t *testing.T) {
	s := Stack{
		StackA: []int{2, 1, 3},
	}

	fmt.Println("Avant Sa:", s.StackA)
	s.Sa()
	fmt.Println("Après Sa:", s.StackA)

	expected := []int{1, 2, 3}

	for i, v := range s.StackA {
		if v != expected[i] {
			t.Errorf("Sa() = %v; want %v", s.StackA, expected)
		}
	}
}

func TestSb(t *testing.T) {
	s := Stack{
		StackB: []int{2, 1, 3},
	}

	fmt.Println("Avant Sb:", s.StackB)
	s.Sb()
	fmt.Println("Après Sb:", s.StackB)

	expected := []int{1, 2, 3}

	for i, v := range s.StackB {
		if v != expected[i] {
			t.Errorf("Sb() = %v; want %v", s.StackB, expected)
		}
	}
}

func TestPb(t *testing.T) {
	s := Stack{
		StackA: []int{1, 2, 3},
		StackB: []int{4, 6, 5},
	}

	fmt.Println("Avant pb stack a:", s.StackA)
	fmt.Println("Avant pb stack b:", s.StackB)
	s.Pb()
	fmt.Println("Après pb stack a:", s.StackA)
	fmt.Println("Après pb stack b:", s.StackB)

	expectedA := []int{2, 3}
	expectedB := []int{1, 4, 6, 5}

	if len(s.StackA) != len(expectedA) {
		t.Errorf("Length of stackA = %v; want %v", len(s.StackA), len(expectedA))
	}

	if len(s.StackB) != len(expectedB) {
		t.Errorf("Length of stackB = %v; want %v", len(s.StackB), len(expectedB))
	}

	for i, v := range s.StackA {
		if v != expectedA[i] {
			t.Errorf("pb() = %v; want %v", s.StackA, expectedA)
		}
	}

	for i, v := range s.StackB {
		if v != expectedB[i] {
			t.Errorf("pb() = %v; want %v", s.StackB, expectedB)
		}
	}
}

func TestPa(t *testing.T) {
	s := Stack{
		StackA: []int{4},
		StackB: []int{1, 2, 3},
	}

	fmt.Println("Avant pa stack a:", s.StackA)
	fmt.Println("Avant pa stack b:", s.StackB)
	s.Pa()
	fmt.Println("Après pa stack a:", s.StackA)
	fmt.Println("Après pa stack b:", s.StackB)

	expectedA := []int{1, 4}
	expectedB := []int{2, 3}

	if len(s.StackA) != len(expectedA) {
		t.Errorf("Length of stackA = %v; want %v", len(s.StackA), len(expectedA))
	}

	if len(s.StackB) != len(expectedB) {
		t.Errorf("Length of stackB = %v; want %v", len(s.StackB), len(expectedB))
	}

	for i := range expectedA {
		if s.StackA[i] != expectedA[i] {
			t.Errorf("pa() = %v; want %v", s.StackA, expectedA)
		}
	}

	for i := range expectedB {
		if s.StackB[i] != expectedB[i] {
			t.Errorf("pa() = %v; want %v", s.StackB, expectedB)
		}
	}
}
