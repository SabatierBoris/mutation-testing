package main

import (
	"fmt"
)

func canBuyAlcool(age int, volume int) bool {
	if age >= 18 && volume <= 10 {
		return true
	}
	return false
}

func main() {
	var age, volume int
	fmt.Print("You want buy alcool, but how old are you ? ")
	_, err := fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Println("That's not a number")
	}
	fmt.Print("How many litter to you want ? ")
	_, err = fmt.Scanf("%d", &volume)
	if err != nil {
		fmt.Println("That's not a number")
	}

	if canBuyAlcool(age, volume) {
		fmt.Println("Enjoy")
	} else {
		fmt.Println("You can't do that")
	}
}
