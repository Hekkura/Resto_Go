package utils

import "fmt"

func MainMenu() {
	fmt.Println("Welcome to the Go Restaurant!")
	fmt.Println("1. View Menu")
	fmt.Println("2. Place Order")
	fmt.Println("3. Exit")
}

func FoodMenu() {
	fmt.Println("Our Menu :")
	foods := InitData()
	for idx, value := range foods {
		fmt.Println(idx+1, ".", value.Food, "- $", value.Price)
	}
	
}

func CloseMenu() {
	fmt.Println("Thank you for dining with us!")
}