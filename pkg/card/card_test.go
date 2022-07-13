package card

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetCard(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	i := 0
	for i <= 15 {
		t.Run(fmt.Sprintf("Lanzada de manos: %d", i), func(t *testing.T) {
			firstCard, secondCard := DealerCards()
			// fmt.Printf("Primera carta: %d \t Segunda carta %d \t", firstCard.Value, secondCard.Value)
			if firstCard.Value >= secondCard.Value {
				t.Errorf("Primera carta es igual o mayor a la segunda. %d vs %d", firstCard.Value, secondCard.Value)
			}

			i++
		})
	}
}

func TestNameCard(t *testing.T) {
	testCard := Card{}
	testCases := []struct {
		value        int
		expectedName string
	}{
		{2, "2"},
		{11, "Jack"},
		{5, "5"},
		{12, "Queen"},
		{13, "King"},
		{14, "Ace"},
	}
	for _, val := range testCases {
		t.Run(fmt.Sprintf("The card display is %s", val.expectedName), func(t *testing.T) {
			convertFigures(figure(val.value), &testCard)

			if testCard.Name == val.expectedName {
				t.Logf("The card name is %s", val.expectedName)
			} else {
				t.Errorf("Value received %s", testCard.Name)
			}
		})
	}
}
