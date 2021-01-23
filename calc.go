package main

import (
	"fmt"
	"math"
)

func main() {
	var choice int

	fmt.Print("[1] Addition\n[2] Subtraction\n[3] Multiplication\n[4] Division\n\n[5] Square Root\n\nWhat would you like to do?\n> ")
	fmt.Scanln(&choice)

	//create two variables two perform math on
	var numberOne float64
	var numberTwo float64

	//bypass for sqrt only needing one argument
	if choice == 5 {
		sqrt()
		return
	}

	//get user input to store into the two variables
	fmt.Print("First number:\n> ")
	fmt.Scanln(&numberOne)
	fmt.Print("Second number:\n> ")
	fmt.Scanln(&numberTwo)
	fmt.Print("\n") // a dumb solution.

	//call the chosen function and insert the two variables into it
	switch choice {
	case 1:
		add(numberOne, numberTwo)
	case 2:
		subt(numberOne, numberTwo)
	case 3:
		mult(numberOne, numberTwo)
	case 4:
		div(numberOne, numberTwo)
	default:
		fmt.Print("\nNot an option!!\n\n") // you may notice i enjoy newlines quite a bit.
		main()
	}
}

//function with two arguments to slot things into
func add(add1 float64, add2 float64) {
	//add the two arguments which now have numberOne and numberTwo stored in them
	addSum := add1 + add2
	fmt.Println(addSum)
}

func subt(subt1 float64, subt2 float64) {
	subtSum := subt1 - subt2
	fmt.Println(subtSum)
}

func mult(mult1 float64, mult2 float64) {
	multSum := mult1 * mult2
	fmt.Println(multSum)
}

func div(div1 float64, div2 float64) {
	divSum := div1 / div2
	fmt.Println(divSum)
}

func sqrt() {
	var sqrtOne float64
	fmt.Print("Number:\n> ")
	fmt.Scanln(&sqrtOne)

	fmt.Println(math.Sqrt(sqrtOne))
}
