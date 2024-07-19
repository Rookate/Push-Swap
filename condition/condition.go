package condition

import (
	"strconv"
)

func IsNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func HasDuplicates(args []string) bool {
	seen := make(map[string]bool)
	for _, arg := range args {
		if seen[arg] {
			return true
		}
		seen[arg] = true
	}
	return false
}
