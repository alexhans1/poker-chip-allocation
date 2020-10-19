package poker

import "testing"

func eval(t *testing.T, winnings map[string]int, expectedWinnings map[string]int) {
	if len(winnings) != len(expectedWinnings) {
		t.Errorf("money allocation is wrong. \nReceived: %+v \nExpected: %+v", winnings, expectedWinnings)
		return
	}
	for pID, winning := range winnings {
		if winning != expectedWinnings[pID] {
			t.Errorf("money allocation is wrong. \nReceived: %+v \nExpected: %+v", winnings, expectedWinnings)
			return
		}
	}
}

func TestOneWinner(t *testing.T) {
	rank := [][]string{{"p2"}, {"p1"}}
	bets := map[string]int{"p1": 5, "p2": 5}
	expectedWinnings := map[string]int{"p1": 0, "p2": 10}

	winnings := Allocate(rank, bets)
	eval(t, winnings, expectedWinnings)
}
func TestDifferentBets(t *testing.T) {
	rank := [][]string{{"p2"}, {"p1"}}
	bets := map[string]int{"p1": 3, "p2": 5}
	expectedWinnings := map[string]int{"p1": 0, "p2": 8}

	winnings := Allocate(rank, bets)
	eval(t, winnings, expectedWinnings)
}
func TestTwoWinners(t *testing.T) {
	rank := [][]string{{"p2", "p3"}, {"p1"}}
	bets := map[string]int{"p1": 10, "p2": 10, "p3": 10}
	expectedWinnings := map[string]int{"p1": 0, "p2": 15, "p3": 15}

	winnings := Allocate(rank, bets)
	eval(t, winnings, expectedWinnings)
}
func TestTwoWinnersWithUnevenWinnings(t *testing.T) {
	rank := [][]string{{"p2", "p3"}, {"p1"}}
	bets := map[string]int{"p1": 11, "p2": 11, "p3": 11}
	expectedWinnings := map[string]int{"p1": 0, "p2": 16, "p3": 16}

	winnings := Allocate(rank, bets)
	eval(t, winnings, expectedWinnings)
}
func TestTwoWinnersWithUnevenWinningsAndBets(t *testing.T) {
	rank := [][]string{{"p2", "p3"}, {"p1"}}
	bets := map[string]int{"p1": 11, "p2": 15, "p3": 9}
	expectedWinnings := map[string]int{"p1": 0, "p2": 21, "p3": 13}

	winnings := Allocate(rank, bets)
	eval(t, winnings, expectedWinnings)
}
func TestTwoWinnersWithLeftOverPot(t *testing.T) {
	rank := [][]string{{"p2", "p3"}, {"p1"}}
	bets := map[string]int{"p1": 17, "p2": 15, "p3": 9}
	expectedWinnings := map[string]int{"p1": 2, "p2": 25, "p3": 13}

	winnings := Allocate(rank, bets)
	eval(t, winnings, expectedWinnings)
}
func TestCascadingAllIn(t *testing.T) {
	rank := [][]string{{"p5"}, {"p4"}, {"p3"}, {"p2"}, {"p1"}}
	bets := map[string]int{"p1": 5, "p2": 4, "p3": 3, "p4": 2, "p5": 1}
	expectedWinnings := map[string]int{"p1": 1, "p2": 2, "p3": 3, "p4": 4, "p5": 5}

	winnings := Allocate(rank, bets)
	eval(t, winnings, expectedWinnings)
}
