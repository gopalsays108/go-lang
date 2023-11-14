package main

import "fmt"

const Token = "here"

func main() {
	var userName string = "Gopal"
	fmt.Println(userName)
	fmt.Printf("VARIABLE IS %T \n\n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type is %T \n\n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable type is %T \n\n", smallValue)

	var smallFloat float64 = 12.2123123123123123123
	fmt.Println(smallFloat)
	fmt.Printf("Variable type is %T \n\n", smallFloat)

	//Default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable type is %T \n\n", anotherVariable)

	//Implicit way to declare variable
	var website = "gopal"

	numberOfDays := 123 //volorous operator
	fmt.Println(numberOfDays)
	fmt.Printf("Variable type is %T \n\n", numberOfDays)
}
