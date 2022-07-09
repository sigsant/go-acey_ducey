package card

import (
	"encoding/json"
	"math/rand"
)

type Card struct {
	Value int
}

func getCard() int {
	return rand.Intn(14) + 1
}

func PlayerCard() *Card {
	playerCard := &Card{}
	playerCard.Value = getCard()
	return playerCard
}

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
			break
		}
	}

	return firstCard, secondCard
}

// ToString convierte el struc en un JSON para su fácil lectura
func (c *Card) ToString() string {
	output, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(output)
}
