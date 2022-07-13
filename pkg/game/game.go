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

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "leyendo desde linea de comandos: ", err)
	}
	return scanner.Text()
}

// showCards muestra las dos cartas del dealer antes de que puedas apostar
func showCards() error {
	dealerCardOne, dealerCardTwo := card.DealerCards()
	fmt.Println("\n\tHere are your next two cards: ")
	fmt.Printf("\t%s", dealerCardOne.Name)
	fmt.Printf("\n\t%s\n", dealerCardTwo.Name)
	return nil
}

func checkActualMoney(bet int) error {
	if actualMoney-bet < 0 {
		actualMoney = 0
		return errors.New("you bet more than your actual money. You have 0 dollars")
	}
	actualMoney -= bet
	return nil
}

func msgBet(hasBet bool, bet int) string {
	if !hasBet {
		return "Chicken"
	} else if hasBet && bet > actualMoney || hasBet && bet < 0 {
		return fmt.Sprintf("You can bet only between 0 and %d", actualMoney)
	}

	//TESTINGME: Update testing with this message
	if err := checkActualMoney(bet); err != nil {
		fmt.Fprintln(os.Stderr, "Error found in your bet: ", err)
	}
	return fmt.Sprintf("You bet %d", bet)
}

// REMEMBERME: checkCondition No integrado
func checkCondition(cardDealerOne, cardDealerTwo, cardPlayer *card.Card) string {
	if cardDealerOne.Value > cardPlayer.Value && cardPlayer.Value < cardDealerTwo.Value {
		return fmt.Sprintf("Congratulations! You have %d dollars", actualMoney)
	}

	return fmt.Sprintf("Sorry. You lose! you have %d dollars", actualMoney)
}

func StartGame() {
	showCards()
	fmt.Printf("\n\tYou have %d dollars\n", actualMoney)
	fmt.Print("\tWhat is your bet? ")
	bet := readInput()

	// TODO: La conversion es redundante en esta parte del código. Pasarla a msgBet
	betToInt, err := strconv.Atoi(bet)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if betToInt <= 0 {
		fmt.Println("\n\t", msgBet(false, betToInt))
	} else {
		fmt.Println("\n\t", msgBet(true, betToInt))
	}
}
