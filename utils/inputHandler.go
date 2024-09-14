package utils

import "fmt"

func InputHandler() string {
	var selector string

	fmt.Printf("=== Please Select an Option : ")
	fmt.Scanf("%s", &selector)

	return selector
}

func InputMenu() (string, int) {
	var name string
	var qty int

	// fmt.Println("Enter the Item name : ")
	// fmt.Scanf("%s", &name)

	// fmt.Println("Enter the quantity : ")
	// fmt.Scanf("%d", &qty)

	fmt.Println("Enter the Item name and quantity : ")
	fmt.Scanln(&name, &qty)

	fmt.Printf("Order Placed : %d %s", qty, name)

	return name, qty

}

func ErrorHandlerInput() {
	fmt.Println("Invalid Input! Please Try Again!")
}