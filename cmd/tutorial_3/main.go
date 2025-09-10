package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello from a function"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Println("Error:", err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of %v / %v is %v\n", numerator, denominator, result)
	// 	return
	// } else {
	// 	fmt.Printf("The result of %v / %v is %v and the remainder is %v\n", numerator, denominator, result, remainder)
	// }

	// break está implícito en cada case
	switch {
	case err != nil:
		fmt.Println("Error:", err.Error())
	case remainder == 0:
		fmt.Printf("The result of %v / %v is %v\n", numerator, denominator, result)
	default:
		fmt.Printf("The result of %v / %v is %v and the remainder is %v\n", numerator, denominator, result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("denominator cannot be zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, nil
}
