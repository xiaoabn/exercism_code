package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	formatID := strings.ReplaceAll(id, " ", "")

	if len(formatID) <= 1 {
		return false
	}

	for _, r := range formatID {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	sum := 0
	for i, r := range formatID {
		d := int(r - '0')
		p := len(formatID) - 1 - i
		if p%2 == 1 {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}
	return sum%10 == 0
}
