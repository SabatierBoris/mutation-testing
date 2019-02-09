package main

func canBuyAlcool(age int, volume int) bool {
	if age >= 18 && volume <= 10 {
		return true
	}
	return false
}
