package main

import (
	"Push-Swap/condition"
	"Push-Swap/stack"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Empty strings")
		return
	}

	args := os.Args[1:]
	if condition.HasDuplicates(args) {
		fmt.Println("Error: There are duplicates in the arguments")
		return
	}

	var s stack.Stack

	tokens := strings.Fields(os.Args[1])
	for _, token := range tokens {
		if !condition.IsNumber(token) {
			fmt.Println("Error: some arguments are not integers")
			return
		}
		num, err := strconv.Atoi(token)
		if err != nil {
			fmt.Println("Error : enable to convert argument to interger")
			return
		}

		s.StackA = append(s.StackA, num)
	}

	//Exemple d'appel de méthode
	s.Ra()
	s.Sa()
	s.Pa()
	s.Pa()
	s.Pa()
	s.Rrb()
	s.Pa()
	s.PrintStack()
}
