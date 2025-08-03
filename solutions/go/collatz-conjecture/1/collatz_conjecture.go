package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	steps := 0
	if n < 1 {
		return n, fmt.Errorf("input can't be 1 or less")
	}
	for n > 1 {
		if n%2 != 0 {
			n = (3 * n) + 1
		} else {
			n = n / 2
		}
		steps++
	}
	return steps, nil
}
