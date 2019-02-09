package main

import "testing"

func TestCanBuyAlcool(t *testing.T) {
	answer := canBuyAlcool(20, 5)
	if answer != true {
		t.Errorf("canBuyAlcool(20, 5) was incorrect, got: %t want: true.", answer)
	}

	answer = canBuyAlcool(10, 5)
	if answer != false {
		t.Errorf("canBuyAlcool(10, 5) was incorrect, got: %t want: false.", answer)
	}
}
