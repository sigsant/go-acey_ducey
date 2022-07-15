package game

import (
	"fmt"
	"testing"
)

func setuptest(t *testing.T) {
	actualMoney = 100
}

func TestBetMessage(t *testing.T) {
	testCases := []struct {
		nameTest    string
		hasBet      bool
		bet         int
		expectedMsg string
	}{
		{"No bet", false, 100, "Chicken"},
		{"Bet with negative number", true, -1, "You can bet only between 0 and 100"},
		{"Bet with 120 dollars", true, 120, "You can bet only between 0 and 100"},
		{"Normal bet", true, 40, "You bet 40"},
	}

	for _, val := range testCases {
		t.Run(fmt.Sprintf(val.nameTest), func(t *testing.T) {
			msg := msgBet(val.hasBet, val.bet)

			if msg != val.expectedMsg {
				t.Error(msg)
			}
		})
	}
}

func TestActualMoney(t *testing.T) {
	setuptest(t)
	moneyCases := []struct {
		money         int
		expectedMoney int
	}{
		{20, 80},
		{10, 70},
		{40, 30},
		{50, 0},
	}

	for _, moneyvalue := range moneyCases {
		_ = checkActualMoney(moneyvalue.money)

		if actualMoney == moneyvalue.expectedMoney {
			t.Logf("Actual money is %d", actualMoney)
		} else {
			t.Errorf("Expected: %d \t Received: %d", moneyvalue.expectedMoney, actualMoney)
		}
	}
}
