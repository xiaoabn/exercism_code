package blackjack

func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

func FirstTurn(card1, card2, dealerCard string) string {
	if card1 == "ace" && card2 == "ace" {
		return "P"
	}
	c1 := ParseCard(card1)
	c2 := ParseCard(card2)
	c3 := ParseCard(dealerCard)

	switch {
	case c1+c2 == 21:
		if c3 != 11 && c3 != 10 {
			return "W"
		}
		return "S"
	case c1+c2 >= 17 && c1+c2 <= 20:
		return "S"
	case c1+c2 >= 12 && c1+c2 <= 16:
		if c3 >= 7 {
			return "H"
		}
		return "S"
	case c1+c2 <= 11:
		return "H"
	}
	return "s"
}
