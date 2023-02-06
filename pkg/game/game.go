package game

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/sigsant/acey_ducey/pkg/card"
)

var (
	actualMoney = 100
)

func init() {
	fmt.Println("\n\tAcey-Ducey is played in the following manner: ")
	fmt.Println("\n\tThe dealer (Computer) deals two cards face up.")
	fmt.Println("\tYou have an option to bet or not to bet depending on whether or not you feel the card will have a value between the  first two.")
	fmt.Println(("\tIf you do not want to bet, input an 0"))

	fmt.Printf("\n\tYou have %d dollars\n", actualMoney)

}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "leyendo desde linea de comandos: ", err)
	}
	return scanner.Text()
}

// showCards shows the two dealer cards before you can bet
func showCards(dealerCardOne, dealerCardTwo card.Card) error {
	// dealerCardOne, dealerCardTwo := card.DealerCards()
	fmt.Printf("\n\tHere are your next two cards:  %s, %s\n", dealerCardOne.Name, dealerCardTwo.Name)
	return nil
}

// checkActualMoney checks if you have enough money to bet
func checkActualMoney(bet int) error {
	if actualMoney-bet < 0 {
		return errors.New("you bet more than your actual money. You have 0 dollars")
	}
	if bet < 0 {
		return errors.New("you can't bet negative numbers")
	}
	actualMoney -= bet
	return nil
}

// msgBet shows the message of the bet
func msgBet(hasBet bool, bet int) string {

	if err := checkActualMoney(bet); err != nil {
		fmt.Fprintln(os.Stderr, "Error found in your bet: ", err)
	}

	if !hasBet {
		return "Chicken"
	} else if hasBet && bet > actualMoney || hasBet && bet < 0 {
		return fmt.Sprintf("You can bet only between 0 and %d", actualMoney)
	} else {
		return fmt.Sprintf("You bet %d", bet)
	}
}

func checkCondition(cardDealerOne, cardDealerTwo, cardPlayer *card.Card) string {
	if cardDealerOne.Value > cardPlayer.Value && cardPlayer.Value < cardDealerTwo.Value {
		return fmt.Sprintf("\tCongratulations! You have %d dollars", actualMoney)
	}
	return fmt.Sprintf("\tSorry. You lose! you have %d dollars", actualMoney)
}

// ReplayGame asks if you want to play again
func ReplayGame() {
	fmt.Print("\tDo you want to play again? (y/n): ")
	replay := readInput()
	if replay == "y" {
		actualMoney = 100
		StartGame()
	} else {
		fmt.Println("\tGame over!")
	}
}

// StartGame starts the game
func StartGame() {
	for actualMoney > 0 {
		// The dealer has two cards
		dealerCardOne, dealerCardTwo := card.DealerCards()
		//The player has a card
		playerCard := card.PlayerCard()

		showCards(*dealerCardOne, *dealerCardTwo)
		fmt.Print("\tWhat is your bet? ")
		bet := readInput()

		// TODO: La conversion es redundante en esta parte del código. Pasarla a msgBet
		betToInt, err := strconv.Atoi(bet)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		if betToInt <= 0 {
			fmt.Println("\n", msgBet(false, betToInt))
		} else {
			fmt.Println("\n", msgBet(true, betToInt))
		}

		fmt.Printf("\tYour card is: %s\n", playerCard.Name)

		// check if the player's card is between the dealer's cards
		fmt.Println(checkCondition(dealerCardOne, dealerCardTwo, playerCard))
	}

	if actualMoney == 0 {
		fmt.Println("\tYou have no more money. Game over!")
		ReplayGame()
	}
}
