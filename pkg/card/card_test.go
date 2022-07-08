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
	for i <= 10 {
		t.Run(fmt.Sprintf("Lanzada de manos: %d", i), func(t *testing.T) {
			firstCard, secondCard := ReturnCard()

			if firstCard.Value < secondCard.Value {
				t.Logf("Primera carta es menor que la segunda. %d vs %d", firstCard.Value, secondCard.Value)
			} else {
				t.Errorf("Primera carta es igual o mayor a la segunda. %d vs %d", firstCard.Value, secondCard.Value)
			}

			i++
		})

	}
}
