package game

import (
	"fmt"
	"testing"
)

func init() {
	actualMoney = 100
}

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

func TestActualMoney(t *testing.T) {
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
