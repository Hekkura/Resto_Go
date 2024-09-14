package main

// import "fmt"
import (
	u "Resto_App/utils"
	"fmt"
)

func main() {

	u.InitData()
	
	for {
		u.MainMenu()

		x := u.InputHandler()
		fmt.Println(x)

		switch x {
			case "1" : u.FoodMenu()
			case "2" : u.InputMenu()
			
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
	