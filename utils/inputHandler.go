package utils

import "fmt"

func InputHandler() string {
	var selector string

	fmt.Printf("=== Please Select an Option : ")
	fmt.Scanln(&selector)

	return selector
}

func InputMenu() (string, int, error) {
	var name string
	var qty int

	fmt.Println("Enter the Item name : ")
	fmt.Scanln(&name)

	fmt.Println("Enter the quantity : ")
	fmt.Scanln(&qty)

	// fmt.Println("Enter the Item name and quantity : ")
	// fmt.Scanln(&name, &qty)


	fmt.Printf("\nOrder Placed : %d %s \n", qty, name)

	return name, qty, nil

}

func ErrorHandlerInput() {
	fmt.Println("Invalid Input! Please Try Again!")
}