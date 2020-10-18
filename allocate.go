package allocate

// Allocate receives a ranking of player ids and a map of [playerId]: betAmount
// It returns a map of [playerId]: receivedAmount
func Allocate(rank [][]string, bets map[string]int) map[string]int {
	maxBettedAmount := 0
	amountLeftInPot := 0
	allocatedWinnings := make(map[string]int)
	for k, v := range bets {
		allocatedWinnings[k] = 0
		amountLeftInPot += v
		if v > maxBettedAmount {
			maxBettedAmount = v
		}
	}

	for _, players := range rank {
		for _, winningPlayerID := range players {
			amountWon := 0
			for playerID, betAmount := range bets {
				amount := betAmount / len(players)
				amountWon += amount
				bets[playerID] -= amount
				allocatedWinnings[winningPlayerID] += amount
				amountLeftInPot -= amount
				if amountLeftInPot == 0 {
					return allocatedWinnings
				}
			}
		}
	}
	return allocatedWinnings
}
