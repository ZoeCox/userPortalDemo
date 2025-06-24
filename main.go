package main

import "fmt"

func main() {
	asciiArt := `   
⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐
 _   _                      
| | | |___  ___ _ __        
| | | / __|/ _ \ '__|       
| |_| \__ \  __/ |          
 \___/|___/\___|_|        _ 
|  _ \ ___  _ __| |_ __ _| |
| |_) / _ \| '__| __/ _  | |
|  __/ (_) | |  | || (_| | |
|_|   \___/|_|   \__\__,_|_| 

⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐⭐`
	fmt.Println(asciiArt)
	var firstName string
	var secondName string
	optionsTxt := "What option would you like to change?\n Change Favourite Color (Type 'color')\n Change Pet (Type 'pet')\n Change Name (Type 'name')"
	fmt.Println("Hello there! Welcome to your user portal - to begin, please confirm your identity -")
	fmt.Println("Enter your first name: ")
	fmt.Scanln(&firstName)
	fmt.Println("Enter your second name: ")
	fmt.Scanln(&secondName)
	fullName := firstName + " " + secondName
	fmt.Println("Hello", fullName, "😊")
	fmt.Println(optionsTxt)
	optionSlice := []string{"No color chosen", "No pet chosen", fullName}

	optionFunc(optionSlice, optionsTxt)
}

func optionFunc(optionSlice []string, optionsTxt string) {
	var viewOptions string
	var option string
	fmt.Println("Type your option here")
	fmt.Scanln(&option)
	switch option {
	case "color":
		fmt.Println("What would you like to change your color too? It is currently:\n", optionSlice[0])
		fmt.Scanln(&optionSlice[0])
		fmt.Println("Your new favourite color is:", optionSlice[0])
	case "pet":
		fmt.Println("What would you like to change your pet too? It is currently:\n", optionSlice[1])
		fmt.Scanln(&optionSlice[1])
		fmt.Println("Your new favourite pet is:", optionSlice[1])
	case "name":
		fmt.Println("What would you like to change your name too? It is currently:\n", optionSlice[2])
		fmt.Scanln(&optionSlice[2])
		fmt.Println("Your new name is:", optionSlice[2])
	default:
		fmt.Println("Invalid option chosen, please review.")
	}
	for i := 0; i < len(optionSlice); i++ {
		println("Number", i+1, "is", optionSlice[i])
	}
	fmt.Println("Type 'view' to see all your options again, or anything else to exit the portal.")
	fmt.Scanln(&viewOptions)
	if viewOptions != "view" {
		fmt.Println("Invalid option chosen, have a good day!")
	} else {
		fmt.Println(optionsTxt)
		optionFunc(optionSlice, optionsTxt)
	}
}
