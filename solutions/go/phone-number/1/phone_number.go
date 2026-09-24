package phonenumber

import (
	"errors"
	"fmt"
	"unicode"
)

func Number(phoneNumber string) (string, error) {
	s := ""
	for _, char := range phoneNumber {
		if unicode.IsDigit(char) {
			s += string(char)
		}
	}

	if len(s) == 10 {
		if s[0]-'0' > 1 && s[3]-'0' > 1 {
			return s, nil
		}
	}

	if len(s) == 11 {
		if s[0]-'0' == 1 && s[1]-'0' > 1 && s[4]-'0' > 1 {
			return s[1:11], nil
		}
	}

	return "phoneNumber", errors.New("invalid phone number")
}

func AreaCode(phoneNumber string) (string, error) {
	s, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.3s", s), nil
}

func Format(phoneNumber string) (string, error) {
	s1, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", s1[:3], s1[3:6], s1[6:10]), nil
}
