package main

import (
	"fmt"
)

func main() {
	var celsius float64
	var choice int

	// Input Celsius temperature
	fmt.Print("Enter temperature in Celsius: ")
	fmt.Scan(&celsius)

	// Menu for conversion choices
	fmt.Println("Select conversion:")
	fmt.Println("1. Fahrenheit")
	fmt.Println("2. Reamur")
	fmt.Println("3. Kelvin")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		// Celsius to Fahrenheit: (C * 9/5) + 32
		fahrenheit := (celsius * 9 / 5) + 32
		fmt.Printf("Temperature in Fahrenheit: %.2f°F\n", fahrenheit)
	case 2:
		// Celsius to Reamur: C * 4/5
		reamur := celsius * 4 / 5
		fmt.Printf("Temperature in Reamur: %.2f°Re\n", reamur)
	case 3:
		// Celsius to Kelvin: C + 273.15
		kelvin := celsius + 273.15
		fmt.Printf("Temperature in Kelvin: %.2fK\n", kelvin)
	default:
		fmt.Println("Invalid choice, please select 1, 2, or 3.")
	}
}
