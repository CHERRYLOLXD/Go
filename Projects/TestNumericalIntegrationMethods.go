package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func integralFunction(value float64) float64 {
	return value * value
}

func getValidInput(prompt string) (float64, error) {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		getValidInput(prompt)
	}
	return input, nil
}

func rectangleMethod(f func(float64) float64, lowerLimit, upperLimit, segmentCount float64) float64 {
	// Implementation of rectangle method in Go
	// ...
	return result
}

func simpsonMethod(f func(float64) float64, lowerLimit, upperLimit, segmentCount float64) float64 {
	// Implementation of simpson method in Go
	// ...
	return result
}

func trapezoidalMethod(f func(float64) float64, lowerLimit, upperLimit, segmentCount float64) float64 {
	// Implementation of trapezoidal method in Go
	// ...
	return result
}

func main() {
	for {
		fmt.Println("Available integration methods:")
		fmt.Println("1 -> Rectangle method")
		fmt.Println("2 -> Simpson method")
		fmt.Println("3 -> Trapezoidal method")

		fmt.Print("Enter your choice (1, 2, or 3): ")
		methodChoice, err := getValidInput("")
		if err != nil || methodChoice < 1 || methodChoice > 3 {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}

		lowerLimit, _ := getValidInput("Enter lower limit:")
		fmt.Println("Enter upper limit:")
		upperLimit, _ := getValidInput("")
		fmt.Println("Enter segment count:")
		segmentCount, _ := getValidInput("")

		var result float64
		switch int(methodChoice) {
		case 1:
			fmt.Println("The rectangle method is selected.")
			result = rectangleMethod(integralFunction, lowerLimit, upperLimit, segmentCount)
		case 2:
			fmt.Println("The simpson method is selected.")
			result = simpsonMethod(integralFunction, lowerLimit, upperLimit, segmentCount)
		case 3:
			fmt.Println("The trapezoidal method is selected.")
			result = trapezoidalMethod(integralFunction, lowerLimit, upperLimit, segmentCount)
		}

		fmt.Printf("Result: %f\n", result)
	}
}
