package main

// import "fmt"
import (
	u "Resto_App/utils"
	"fmt"
)


type Customer interface { 
	//Customer : bisa View menu, Place Order, Exit
}

type Staff interface { 
	//Owner : View Menu, Place Order, Add new menu, Edit menu, Delete Menu, Exit
}

func main() {
	fmt.Println("=== Resto_GO v 0.5 ===")
	fmt.Println("By Henry H. Prasetya")
	// u.InitData()
	
	for {
		
		u.MainMenu()

		x := u.InputHandler()
		fmt.Println(x)

		switch x {
			case "1" : 	u.FoodMenu()
			case "2" : 	u.InputMenu()
			case "3" :  u.CloseMenu()
			
			default : 	
				u.ErrorHandlerInput()
		}
		//Exit handler
		if x == "3" {
			break
		}
	}




	
	
}
	