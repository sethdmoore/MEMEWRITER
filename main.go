package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var input []string
	// take all arguments (except binary name)
	input = os.Args[1:]
	// Convert os.Args []string slice to string
	mergedArgs := strings.Join(input, "")
	// Split string into []string{"s", "l", "i", "c", "e"}
	inputSlice := strings.Split(mergedArgs, "")

	for i := 0; i < len(mergedArgs); i++ {
		// Join string "foo" into "f o o"
		fmt.Printf("%s\n", strings.ToUpper(strings.Join(inputSlice, " ")))
		// Take the first element and all remaining elements
		firstElement, lastElements := inputSlice[0], inputSlice[1:]
		// rearrange the slice
		inputSlice = append(lastElements, firstElement)
	}
}
