package main

import "testing"

func TestCanBuyAlcool(t *testing.T) {
	testCases := []struct {
		age    int
		volume int
		answer bool
	}{
		{10, 2, false},
		{20, 2, true},
	}

	for _, testCase := range testCases {
		answer := canBuyAlcool(testCase.age, testCase.volume)
		if answer != testCase.answer {
			t.Errorf("canBuyAlcool(%d, %d) was incorrect, got: %t want: %t.", testCase.age, testCase.volume, answer, testCase.answer)
		}
	}
}
