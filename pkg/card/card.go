package card

import (
	"encoding/json"
	"math/rand"
	"strconv"
)

type Card struct {
	Value int
	Name  string
}

type figure int

const (
	jack = figure(iota + 11)
	queen
	king
	ace
)

// getCard generate a range of cards from 2 to 14
// 11 (Jack), 12 (Queen), 13 (King), 14 (Ace)
func getCard() int {
	return rand.Intn(14) + 1
}

func PlayerCard() *Card {
	playerCard := &Card{}
	playerCard.Value = getCard()
	return playerCard
}

// Dealercards devuelve las dos primeras cartas del dealer
func DealerCards() (*Card, *Card) {
	firstCard := &Card{}
	secondCard := &Card{}

	firstCard.Value = getCard()
	secondCard.Value = getCard()

	for {
		if firstCard.Value >= secondCard.Value || firstCard.Value == secondCard.Value {
			firstCard.Value = getCard()
			secondCard.Value = getCard()
		} else {
			convertFigures(figure(firstCard.Value), firstCard)
			convertFigures(figure(secondCard.Value), secondCard)
			break
		}
	}

	return firstCard, secondCard
}

// convertFigures añade el nombre de la figura a la carta segun el numero generado
func convertFigures(f figure, c *Card) {
	switch f {
	case jack:
		c.Name = "Jack"
	case queen:
		c.Name = "Queen"
	case king:
		c.Name = "King"
	case ace:
		c.Name = "Ace"
	default:
		c.Name = strconv.Itoa(int(f))
	}
}

// ToString convierte el struc en un JSON para su fácil lectura
// Empleado para testing
func (c *Card) ToString() string {
	output, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(output)
}
