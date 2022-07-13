package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/sigsant/acey_ducey/pkg/card"
)

const (
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
	fmt.Printf("\n\tYou have %d dollars\n", actualMoney)
	fmt.Print("\tWhat is your bet? ")
	bet := readInput()
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
