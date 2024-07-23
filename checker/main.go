package main

import (
	"Push-Swap/algorithms"
	"Push-Swap/instructions"
	"Push-Swap/stack"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Parse command line arguments
	flag.Parse()
	args := flag.Args()

	// Check if there are no arguments
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No arguments provided")
		os.Exit(1)
	}

	// The first argument is the initial stack (formatted as a space-separated list of integers)
	initialStackStr := args[0]

	// Convert the initial stack argument to a slice of integers
	stackA := parseStack(initialStackStr)
	if stackA == nil {
		fmt.Fprintln(os.Stderr, "Error: Invalid stack format")
		os.Exit(1)
	}

	// Initialize StackA and StackB directly
	s := &stack.Stack{
		StackA: stackA,
		StackB: []int{}, // Initialize StackB as empty
	}

	// Determine whether to read instructions from stdin or from the remaining command line arguments
	var instructionsArgs []string
	if len(args) > 1 {
		// Instructions provided as command line arguments
		instructionsArgs = args[1:]
	} else {
		// Read instructions from stdin
		instructionsArgs = readInstructionsFromStdin()
	}

	// Execute each instruction
	for _, instruction := range instructionsArgs {
		instructions.ExecuteInstruction(s, instruction)
	}

	// Check if stackA is sorted and stackB is empty
	if algorithms.IsSorted(s.StackA) && len(s.StackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

// Convert a space-separated string of numbers into a slice of integers
func parseStack(input string) []int {
	parts := strings.Fields(input)
	stack := make([]int, len(parts))

	for i, part := range parts {
		var num int
		if _, err := fmt.Sscanf(part, "%d", &num); err != nil {
			return nil // Return nil if parsing fails
		}
		stack[i] = num
	}
	return stack
}

// Read instructions from stdin
func readInstructionsFromStdin() []string {
	var instructions []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			instructions = append(instructions, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading instructions from stdin:", err)
		os.Exit(1)
	}
	return instructions
}
