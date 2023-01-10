package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// ask user if they want to add, subtract, multiply, or divide
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What would you like to do? (add, subtract, multiply, divide)")
	input, _ := reader.ReadString('\n')

	// ask user for first number
	fmt.Println("What is the first number?")
	fNum, _ := reader.ReadString('\n')

	// try to convert first number to float64
	// remove the newline character
	floatFNum, err := strconv.ParseFloat(fNum[:len(fNum)-1], 64)
	if err != nil {
		panic(err)
	}

	// ask user for second number
	fmt.Println("What is the second number?")
	sNum, _ := reader.ReadString('\n')

	// try to convert second number to float64
	floatSNum, err := strconv.ParseFloat(sNum[:len(sNum)-1], 64)
	if err != nil {
		panic(err)
	}

	// do the math
	switch input[:len(input)-1] {
	case "add":
		fmt.Println(floatFNum + floatSNum)
	case "subtract":
		fmt.Println(floatFNum - floatSNum)
	case "multiply":
		fmt.Println(floatFNum * floatSNum)
	case "divide":
		fmt.Println(floatFNum / floatSNum)
	default:
		fmt.Println("Invalid input")
	}

}
