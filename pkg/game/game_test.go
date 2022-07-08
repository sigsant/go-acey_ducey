package game

import (
	"fmt"
	"testing"
)

func TestBetMessage(t *testing.T) {
	testCases := []struct {
		hasBet      bool
		bet         int
		expectedMsg string
	}{
		{false, 100, "Chicken"},
		{true, -1, "You can bet only between 0 and 100"},
		{true, 120, "You can bet only between 0 and 100"},
		{true, 40, "You bet 40"},
	}

	for _, val := range testCases {
		t.Run(fmt.Sprintf("Apuesta %t realizada con %d", val.hasBet, val.bet), func(t *testing.T) {
			msg := msgBet(val.hasBet, val.bet)

			if msg == val.expectedMsg {
				t.Log(val.expectedMsg)
			} else {
				t.Error(msg)
			}
		})
	}
}
