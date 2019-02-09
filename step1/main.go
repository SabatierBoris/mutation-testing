package main

import (
	"fmt"
)

func canBuyAlcool(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}

func main() {
	var age int
	fmt.Print("You want buy alcool, but how old are you ? ")
	_, err := fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Println("That's not a number")
	}

	if canBuyAlcool(age) {
		fmt.Println("Enjoy")
	} else {
		fmt.Println("No way, you're too young !!")
	}
}
