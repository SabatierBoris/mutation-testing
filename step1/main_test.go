package main

import "testing"

func TestCanBuyAlcool(t *testing.T) {
	anwser := canBuyAlcool(20)
	if anwser != true {
		t.Errorf("canBuyAlcool(20) was incorrect, got: %t want: true.", anwser)
	}
}
