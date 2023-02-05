package game

import (
	"fmt"
	"testing"
)

func setuptest(t *testing.T) {
	actualMoney = 100
}

// Test the message of the bet
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

// Test the amount of money after the bet
func TestActualMoney(t *testing.T) {
	setuptest(t)
	moneyCases := []struct {
		bet           int
		expectedMoney int
	}{
		{20, 80},
		{10, 70},
		{40, 30},
		{50, 30}, // This is the limit of money. If bet more than the limit, the money is not updated
	}

	for _, moneyvalue := range moneyCases {
		t.Run(fmt.Sprintf("Bet %d", moneyvalue.bet), func(t *testing.T) {
			_ = checkActualMoney(moneyvalue.bet)

			if actualMoney == moneyvalue.expectedMoney {
				t.Logf("Actual money is %d", actualMoney)
			} else {
				t.Errorf("Expected: %d \t Received: %d", moneyvalue.expectedMoney, actualMoney)
			}
		})
	}
}

// Test the errors of the bet
func TestBetError(t *testing.T) {
	setuptest(t)
	errorCases := []struct {
		bet         int
		expectedErr error
	}{
		{101, fmt.Errorf("you bet more than your actual money. You have 0 dollars")},
		{-1, fmt.Errorf("you can't bet negative numbers")},
	}

	for _, errorvalue := range errorCases {
		t.Run(fmt.Sprintf("Bet %d", errorvalue.bet), func(t *testing.T) {
			err := checkActualMoney(errorvalue.bet)

			if err != nil {
				t.Logf("Error found in your bet: %s", err)
			} else {
				t.Errorf("Expected: %s \t Received: %s", errorvalue.expectedErr, err)
			}
		})
	}
}
