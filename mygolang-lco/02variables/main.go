package main

import "fmt"

const LoginToken string = "qwerty" // Public 

func main() {
	var username string = "Piyush"
	fmt.Println(username)
	fmt.Printf("variabe is of Type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variabe is of Type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variabe is of Type: %T \n", smallVal)

	var smallFloat float64 = 255.123456789101112131415
	fmt.Println(smallFloat)
	fmt.Printf("variabe is of Type: %T \n", smallFloat)

	//default values and some aliases
	var intDval int
	fmt.Println(intDval)

	var strDval string
	fmt.Println(strDval)

	// implecit type
	var web = "Piyush"
	fmt.Println(web)

	// no var Style
	jwtToken := 3000
	fmt.Println(jwtToken)

	fmt.Println(LoginToken)

}
