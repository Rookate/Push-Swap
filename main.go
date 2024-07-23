package main

import (
	algorithm "Push-Swap/algorithms"
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

	args := os.Args[1]
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
	algorithm.Resolve(&s)
	result := algorithm.Checker(&s)
	fmt.Println("Stack A:", s.StackA)
	fmt.Println("Stack B:", s.StackB)
	fmt.Println("Result:", result)
	fmt.Println("Instruction Count:", s.InstructionCount)
	fmt.Println("Flags:", strings.Join(s.Operation, "\\n"))
}
