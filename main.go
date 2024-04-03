package main

import (
	"fmt"
)

func main() {
	var coffeeOption int16
	var value float32

	fmt.Println("Welcome to the coffee shop!")
	fmt.Println("We have these coffee options:")
	fmt.Println("1. Espresso")
	fmt.Println("2. Latte")
	fmt.Println("3. Cappuccino")
	fmt.Println("4. Americano")
	fmt.Print("Please choose a coffee option: ")
	_, err := fmt.Scanf("%d", &coffeeOption)

	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	orderCoffee(coffeeOption)

	fmt.Println("You can pay with $0.25, $0.50 or $1.00")
	fmt.Printf("Please enter your payment amount: ")

	_, err = fmt.Scanf("%f", &value)

	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	if value == 0 || value > 1 {
		fmt.Println("Invalid value.")
	}

	payment(value, coffeeOption)
}

func orderCoffee(coffeeOption int16) {
	switch coffeeOption {
	case 1:
		fmt.Println("You ordered an espresso, it's $1.")
	case 2:
		fmt.Println("You ordered a latte, it's $1.")
	case 3:
		fmt.Println("You ordered a cappuccino, it's $2.")
	case 4:
		fmt.Println("You ordered an americano, it's $2.")
	default:
		fmt.Println("Invalid option.")
	}
}

func payment(value float32, coffeeOption int16) {
	var price float32
	if coffeeOption == 3 || coffeeOption == 4 {
		price = 2.0
	} else {
		price = 1.0
	}

	var total float32 = 0.0

	for {
		total += value

		if total >= price {
			break
		}

		fmt.Println("You can pay with $0.25, $0.50 and $1.00")
		fmt.Printf("Please enter your payment amount: ")
		_, err := fmt.Scanf("%f", &value)
		if err != nil {
			fmt.Println("Invalid input.")
			return
		}
	}

	change := total - price
	fmt.Printf("Your total is $%.2f\n", total)
	fmt.Printf("You paid $%.2f\n", price)
	if change > 0 {
		fmt.Printf("Your change is $%.2f\n", change)
	}
}
