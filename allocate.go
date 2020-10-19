package poker

import (
	"github.com/thoas/go-funk"
)

type PlayerMap map[string]int

// Allocate receives a ranking of player ids and a map of [playerId]: betAmount
// It returns a map of [playerId]: receivedAmount
func Allocate(rank [][]string, b PlayerMap) PlayerMap {
	bets := make(PlayerMap)
	for k, v := range b {
		bets[k] = v
	}
	amountLeftInPot := 0
	allocatedWinnings := make(map[string]int)
	for k, v := range bets {
		allocatedWinnings[k] = 0
		amountLeftInPot += v
	}

	for amountLeftInPot > 0 {
		amountLeftInPot -= getSidePot(getWinners(rank, (bets)), (bets), allocatedWinnings)
	}
	return allocatedWinnings

}

func getSidePot(winners []string, bets PlayerMap, allocatedWinnings PlayerMap) int {
	minWinningBet := bets[winners[0]]
	for _, pID := range winners {
		minWinningBet = minInt(minWinningBet, bets[pID])
	}

	sidePot := 0
	for pID, betAmount := range bets {
		a := minInt(minWinningBet, betAmount)
		sidePot += a
		bets[pID] -= a
	}

	for _, winningPlayerID := range winners {
		allocatedWinnings[winningPlayerID] += sidePot / len(winners)
	}

	return sidePot
}

func getWinners(rank [][]string, bets PlayerMap) []string {
	for _, players := range rank {
		winners := funk.FilterString(players, func(pID string) bool {
			return bets[pID] > 0
		})
		if len(winners) > 0 {
			return winners
		}
	}
	return []string{}
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}
