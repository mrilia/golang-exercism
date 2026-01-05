package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card{
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

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    playerValue := ParseCard(card1)+ParseCard(card2)
    
	switch {
        case isAce(card1) && isAce(card2):
        	return "P"
        case playerValue == 21 && !isAceFaceOrTen(dealerCard):
    		return "W"
        case playerValue == 21:
        	return "S"
        case playerValue >= 17 && playerValue <= 20:
			return "S"
        case playerValue >= 12 && playerValue <= 16 && ParseCard(dealerCard)>=7:
        	return "H"
        case playerValue >= 12 && playerValue <= 16:
        	return "S"
        case playerValue <= 11:
		return "H"
		default:
			return "H" 
    }
}

func isAceFaceOrTen(card string) bool {
	return card == "ace" ||
		card == "jack" ||
		card == "queen" ||
		card == "king" ||
		card == "ten"
}

// isAce checks if a card is an ace
func isAce(card string) bool {
	return card == "ace"
}