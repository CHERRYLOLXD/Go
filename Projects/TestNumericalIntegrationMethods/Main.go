package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getValidInput(prompt string, typeValue any) any {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	switch typeValue.(type) {
	case uint8:
		number, parseError := strconv.ParseUint(input, 10, 8)
		if parseError != nil {
			fmt.Println("Invalid input. Please enter a valid |uint8|.")
			return getValidInput(prompt, typeValue)
		}
		return uint8(number)
	case uint64:
		number, parseError := strconv.ParseUint(input, 10, 64)
		if parseError != nil {
			fmt.Println("Invalid input. Please enter a valid |uint64|.")
			return getValidInput(prompt, typeValue)
		}
		return uint64(number)
	case float64:
		number, parseError := strconv.ParseFloat(input, 64)
		if parseError != nil {
			fmt.Println("Invalid input. Please enter a valid |float64|.")
			return getValidInput(prompt, typeValue)
		}
		return float64(number)
	default:
		fmt.Println("Unsupported data type")
		return nil
	}
}

func rectangleMethod(integralFunction func(float64) float64, lowerLimit, upperLimit float64, segmentCount uint64) (float64, error) {
	if segmentCount <= 0 {
		return 0, errors.New("Segment count must be greater than zero.")
	}
	segmentWidth := (upperLimit - lowerLimit) / float64(segmentCount)
	sum := 0.0
	for segmentIndex := uint64(0); segmentIndex < segmentCount; segmentIndex++ {
		midpoint := lowerLimit + float64(segmentIndex) * segmentWidth + 0.5 * segmentWidth
		sum += integralFunction(midpoint)
	}
	return segmentWidth * sum, nil
}

func simpsonMethod(integralFunction func(float64) float64, lowerLimit, upperLimit float64, segmentCount uint64) (float64, error) {
	if segmentCount <= 0 {
		return 0, errors.New("Segment count must be greater than zero.")
	}
	segmentWidth := (upperLimit - lowerLimit) / float64(segmentCount)
	sum := integralFunction(lowerLimit) + integralFunction(upperLimit)
	for segmentIndex := uint64(1); segmentIndex < segmentCount; segmentIndex++ {
		x := lowerLimit + float64(segmentIndex) * segmentWidth
		weight := 4.0
		if segmentIndex % 2 == 0 {
			weight = 2.0
		}
		sum += weight * integralFunction(x)
	}
	return (segmentWidth / 3.0) * sum, nil
}

func trapezoidalMethod(integralFunction func(float64) float64, lowerLimit, upperLimit float64, segmentCount uint64) (float64, error) {
	if segmentCount <= 0 {
		return 0, errors.New("Segment count must be greater than zero.")
	}
	segmentWidth := (upperLimit - lowerLimit) / float64(segmentCount)
	sum := integralFunction(lowerLimit) + integralFunction(upperLimit)
	for segmentIndex := uint64(1); segmentIndex < segmentCount; segmentIndex++ {
		x := lowerLimit + float64(segmentIndex) * segmentWidth
		sum += 2.0 * integralFunction(x)
	}
	return (segmentWidth / 2.0) * sum, nil
}

func integralFunction(value float64) float64 {
	return value * value
}

func main() {
	for {
		fmt.Println("Available integration methods:")
		fmt.Println("1 -> Rectangle method")
		fmt.Println("2 -> Simpson method")
		fmt.Println("3 -> Trapezoidal method")
		methodChoice := getValidInput("Enter your choice (1, 2, or 3):", uint8(0)).(uint8)
		switch methodChoice {
		case 1:
			fmt.Println("The rectangle method is selected.")
			lowerLimit := getValidInput("Enter lower limit:", float64(0)).(float64)
			upperLimit := getValidInput("Enter upper limit:", float64(0)).(float64)
			segmentCount := getValidInput("Enter segment count:", uint64(0)).(uint64)
			result, rectangleError := rectangleMethod(integralFunction, lowerLimit, upperLimit, segmentCount)
			if rectangleError != nil {
				fmt.Println(rectangleError)
				break
			}
			fmt.Printf("Result using the rectangle method: %f\n", result)
		case 2:
			fmt.Println("The simpson method is selected.")
			lowerLimit := getValidInput("Enter lower limit:", float64(0)).(float64)
			upperLimit := getValidInput("Enter upper limit:", float64(0)).(float64)
			segmentCount := getValidInput("Enter segment count:", uint64(0)).(uint64)
			result, simpsonError := simpsonMethod(integralFunction, lowerLimit, upperLimit, segmentCount)
			if simpsonError != nil {
				fmt.Println(simpsonError)
				break
			}
			fmt.Printf("Result using the simpson method: %f\n", result)
		case 3:
			fmt.Println("The trapezoidal method is selected.")
			lowerLimit := getValidInput("Enter lower limit:", float64(0)).(float64)
			upperLimit := getValidInput("Enter upper limit:", float64(0)).(float64)
			segmentCount := getValidInput("Enter segment count:", uint64(0)).(uint64)
			result, trapezoidalError := trapezoidalMethod(integralFunction, lowerLimit, upperLimit, segmentCount)
			if trapezoidalError != nil {
				fmt.Println(trapezoidalError)
				break
			}
			fmt.Printf("Result using the trapezoidal method: %f\n", result)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
