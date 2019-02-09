package main

import "testing"

func TestCanBuyAlcool(t *testing.T) {
	anwser := canBuyAlcool(20, 5)
	if anwser != true {
		t.Errorf("canBuyAlcool(20, 5) was incorrect, got: %t want: true.", anwser)
	}

	anwser = canBuyAlcool(10, 5)
	if anwser != false {
		t.Errorf("canBuyAlcool(10, 5) was incorrect, got: %t want: false.", anwser)
	}
}
