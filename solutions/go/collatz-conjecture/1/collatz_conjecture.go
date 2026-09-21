package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
    sum := 0
	if n >= 1 {
		for n > 1 {
			if n%2 == 0 {
				n /= 2
			} else {
				n = 3*n + 1
			}
			sum++
		}
		return sum, nil
	}
	return n, errors.New("is not a correct number")
}
