package main

import "testing"

func TestCanBuyAlcool(t *testing.T) {
	answer := canBuyAlcool(20)
	if answer != true {
		t.Errorf("canBuyAlcool(20) was incorrect, got: %t want: true.", answer)
	}
}
