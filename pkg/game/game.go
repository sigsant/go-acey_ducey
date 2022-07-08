package game

import "fmt"

const (
	actualMoney = 100
)

func msgBet(hasBet bool, bet int) string {
	if !hasBet {
		return "Chicken"
	} else if hasBet && bet > actualMoney || hasBet && bet < 0 {
		return fmt.Sprintf("You can bet only between 0 and %d", actualMoney)
	}

	return fmt.Sprintf("You bet %d", bet)
}
