package techpalace

import (
	"fmt"
	"strings"
)

func WelcomeMessage(customer string) string {
    customer = strings.ToUpper(customer)
	return fmt.Sprintf("Welcome to the Tech Palace, %s", customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	star := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%s\n%s\n%s", star, welcomeMsg, star)
}

func CleanupMessage(oldMsg string) string {
	oldMsg = strings.ReplaceAll(oldMsg, "*", "")
	oldMsg = strings.ReplaceAll(oldMsg, "\n", "")
	oldMsg = strings.Trim(oldMsg, " ")
	return oldMsg
}
