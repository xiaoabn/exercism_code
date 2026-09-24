package phonenumber

import (
	"errors"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	var b strings.Builder
	for _, char := range phoneNumber {
		if char >= '0' && char <= '9' {
			b.WriteRune(char)
		}
	}
	s := b.String()

	if len(s) == 11 && s[0] == '1' {
		s = s[1:]
	}
	if len(s) != 10 {
		return "", errors.New("invalid phone number")
	}

	if s[0] < '2' || s[3] < '2' {
		return "", errors.New("invalid phone number")
	}

	return s, nil
}

func AreaCode(phoneNumber string) (string, error) {
	s, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return s[:3], nil
}

func Format(phoneNumber string) (string, error) {
	s, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return "(" + s[:3] + ") " + s[3:6] + "-" + s[6:], nil
}
