package main

import "fmt"

func main() {
	var choice int
	fmt.Print("[1] Addition\n[2] Subtraction\n[3] Multiplication\n[4] Division\n\nWhat would you like to do?\n> ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		add()
	case 2:
		subt()
	case 3:
		mult()
	case 4:
		div()
	default:
		fmt.Print("\nNot an option!!\n\n") // you may notice i enjoy newlines quite a bit.
		main()
	}
}

func rerun() {
	var choice string
	fmt.Println("Is there anything else you would like to do? [y/n]")
	fmt.Scanln(&choice)

	switch choice {
	case "y":
		main()
	case "n":
		fmt.Println("Bye")
	default:
		fmt.Println("Not an option!!")
		rerun()
	}

}

func add() {
	var addOne float64
	var addTwo float64

	fmt.Print("First number:\n> ")
	fmt.Scanln(&addOne)
	fmt.Print("Second number:\n> ")
	fmt.Scanln(&addTwo)
	fmt.Print("\n") // a dumb solution.

	fmt.Println(addOne + addTwo)
	rerun()
}

func subt() {
	var subtOne float64
	var subtTwo float64

	fmt.Print("First number:\n> ")
	fmt.Scanln(&subtOne)
	fmt.Print("Second number:\n> ")
	fmt.Scanln(&subtTwo)
	fmt.Print("\n")

	fmt.Println(subtOne - subtTwo)
	rerun()
}

func mult() {
	var multOne float64
	var multTwo float64

	fmt.Print("First number:\n> ")
	fmt.Scanln(&multOne)
	fmt.Print("Second number:\n> ")
	fmt.Scanln(&multTwo)
	fmt.Print("\n")

	fmt.Println(multOne * multTwo)
	rerun()
}

func div() {
	var divOne float64
	var divTwo float64

	fmt.Print("First number:\n> ")
	fmt.Scanln(&divOne)
	fmt.Print("Second number:\n> ")
	fmt.Scanln(&divTwo)
	fmt.Print("\n")

	fmt.Println(divOne / divTwo)
	rerun()
}
