package condition

import (
	"strconv"
	"strings"
)

func IsNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func HasDuplicates(args string) bool {
	elements := strings.Fields(args)
	seen := make(map[string]struct{})

	for _, element := range elements {
		if _, exists := seen[element]; exists {
			return true
		}
		seen[element] = struct{}{}
	}
	return false
}
