package main

import (
	"fmt"

	"github.com/Zaressa/pivottech/calculator"
)

func main() {
	fmt.Println(calculator.Add(2, 2))
	fmt.Println(calculator.Subtract(4, 1))
	fmt.Println(calculator.Multiply(2, 2))
	fmt.Println(calculator.Divide(4, 2))
	fmt.Println(calculator.Divide(4, 0))
}
