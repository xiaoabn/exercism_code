package hamming

import (
    "errors"
)

func Distance(a, b string) (int, error) {
    num := 0
	if len(a) != len(b) {
		return 0, errors.New("lengths are different")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			num++
		}
	}
	return num, nil
}
