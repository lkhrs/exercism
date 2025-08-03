package twofer

import "fmt"

// function ShareWith takes a name and returns a string containing the name, or returns a different string if the name is empty.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
