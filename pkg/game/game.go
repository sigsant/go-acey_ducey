package game

import (
	"fmt"

	"github.com/sigsant/acey_ducey/pkg/card"
)

const (
	actualMoney = 100
)

// showCards muestra las dos cartas del dealer antes de que puedas apostar
func showCards() error {
	dealerCardOne, dealerCardTwo := card.DealerCards()
	fmt.Println("\n\tHere are your next two cards: ")
	fmt.Printf("\t%d", dealerCardOne.Value)
	fmt.Printf("\n\t%d\n", dealerCardTwo.Value)
	return nil
}

func msgBet(hasBet bool, bet int) string {
	if !hasBet {
		return "Chicken"
	} else if hasBet && bet > actualMoney || hasBet && bet < 0 {
		return fmt.Sprintf("You can bet only between 0 and %d", actualMoney)
	}

	return fmt.Sprintf("You bet %d", bet)
}

func StartGame() {
	showCards()
	fmt.Println("\n\tWhat is your bet?")
	return
}
