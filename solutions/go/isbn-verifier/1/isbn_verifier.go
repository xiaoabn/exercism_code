package isbnverifier

import (
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < 10; i++ {
		ch := isbn[i]
		var d int

		switch {
		case ch >= '0' && ch <= '9':
			d = int(ch - '0')
		case ch == 'X' && i == 9:
			d = 10
		default:
			return false
		}
		sum += d * (10 - i)
	}

	return sum%11 == 0
}
