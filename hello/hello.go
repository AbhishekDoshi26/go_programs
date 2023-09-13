package hello

import "strings"

// This uses array of strings
func Say(names []string) string {
	if len(names) == 0 {
		// If there are no names, it adds world to the array.
		names = []string{"world"}
	}
	return "Hello " + strings.Join(names, ", ")
}
