package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	for {

		varGreetings := greetings()

		getOperation(varGreetings)

	}

}

func add() {

	var numbersToAdd []int
	var result int

	fmt.Println("########################")
	fmt.Println("Provide the list of numbers to add:")
	for {
		var number int
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println("You have entered a non integer as the last number, the addition result is displayed below.")
			break
		}
		numbersToAdd = append(numbersToAdd, number)
	}

	for _, digits := range numbersToAdd {
		result = result + digits
	}
	fmt.Printf("You have provided this list: %v.\n", numbersToAdd)
	fmt.Printf("The sum of the numbers in the provided list is: %v \n", result)

}

func subtract() {

	var firstNum float64
	var secondNum float64

	fmt.Println("########################")
	fmt.Println("Provide the first number")
	fmt.Scan(&firstNum)

	fmt.Println("Provide the second number")
	fmt.Scan(&secondNum)

	result := firstNum - secondNum

	fmt.Println("########################")
	fmt.Printf("The result of subtracting the two numbers provided is: %v\n", result)

}

func multiply() {

	var numbersToMultiply []float64
	result := float64(1)

	fmt.Println("########################")
	fmt.Println("Provide the list of numbers to multiply:")
	for {
		var number float64
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println("You have entered a non integer as the last number, the multiplication result is displayed below ")
			break
		}
		numbersToMultiply = append(numbersToMultiply, number)
	}

	for _, digits := range numbersToMultiply {
		result = result * digits
	}
	fmt.Printf("You have provided this list: %v.\n", numbersToMultiply)
	fmt.Printf("The multiplication of the numbers in the provided list is: %v \n", result)

}

func divide() {

	var firstNum float64
	var secondNum float64

	var result float64

	fmt.Println("########################")
	fmt.Println("Provide the first number")
	fmt.Scan(&firstNum)

	fmt.Println("Provide the second number")
	fmt.Scan(&secondNum)

	if secondNum == 0 {
		fmt.Println("You cannot divide a number by 0")
		return
	} else {
		result = firstNum / secondNum
	}

	fmt.Println("########################")
	fmt.Printf("The result of dividing the two numbers provided is: %v\n", result)

}

func exponent() {

	var baseNumber float64
	var exponent float64
	var result float64

	fmt.Println("########################")
	fmt.Println("Provide the base number")
	fmt.Scan(&baseNumber)

	fmt.Println("Provide the exponential number")
	fmt.Scan(&exponent)

	result = math.Pow(baseNumber, exponent)

	fmt.Println("########################")
	fmt.Printf("The result for %v raised to power %v is: %v  \n", baseNumber, exponent, result)

}

func squart() {
	var wholeNumber float64
	var result float64

	fmt.Println("########################")
	fmt.Println("Provide the number to get the Square root: ")
	fmt.Scan(&wholeNumber)

	result = math.Sqrt(wholeNumber)

	fmt.Println("########################")
	fmt.Printf("The square root of %v is : %v  \n", wholeNumber, result)

}

func greetings() string {

	var userOperation string

	fmt.Println("########################")
	fmt.Println("     ")
	fmt.Println("Welcome to using Habeeb's cli-calculator")
	fmt.Println("The following Operations are supported: Addition, Subtraction, Multiplication, Division, Exponentiation and Squareroot")
	fmt.Println("     ")
	fmt.Println("########################")
	fmt.Println("State the Operation you want to perform:")

	fmt.Scanln(&userOperation)

	return userOperation
}

func getOperation(operation string) {

	switch operation {
	case "Addition":
		add()
	case "Subtraction":
		subtract()
	case "Multiplication":
		multiply()
	case "Division":
		divide()
	case "Exponentiation":
		exponent()
	case "Squareroot":
		squart()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("No other operations left to perform")
	}

}
