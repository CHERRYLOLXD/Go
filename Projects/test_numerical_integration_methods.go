package main

import (
	"fmt"
	"test_numerical_integration_methods/Packages/Input/Validation"
	"test_numerical_integration_methods/Packages/NumericalIntegration/RectangleMethod"
	"test_numerical_integration_methods/Packages/NumericalIntegration/SimpsonMethod"
	"test_numerical_integration_methods/Packages/NumericalIntegration/TrapezoidalMethod"
)

func integralFunction(value float64) float64 {
	return value * value
}

func main() {
	for {
		fmt.Println("Available integration methods:")
		fmt.Println("1 -> Rectangle method")
		fmt.Println("2 -> Simpson method")
		fmt.Println("3 -> Trapezoidal method")
		methodChoice := input_validation.getValidInput("Enter your choice (1, 2, or 3):", uint8(0)).(uint8)
		switch methodChoice {
		case 1:
			fmt.Println("The rectangle method is selected.")
			lowerLimit := input_validation.getValidInput("Enter lower limit:", float64(0)).(float64)
			upperLimit := input_validation.getValidInput("Enter upper limit:", float64(0)).(float64)
			segmentCount := input_validation.getValidInput("Enter segment count:", uint64(0)).(uint64)
			fmt.Printf("Result using the rectangle method: %f\n", rectangle_method.rectangleMethod(integralFunction, lowerLimit, upperLimit, segmentCount))
		case 2:
			fmt.Println("The simpson method is selected.")
			lowerLimit := input_validation.getValidInput("Enter lower limit:", float64(0)).(float64)
			upperLimit := input_validation.getValidInput("Enter upper limit:", float64(0)).(float64)
			segmentCount := input_validation.getValidInput("Enter segment count:", uint64(0)).(uint64)
			fmt.Printf("Result using the simpson method: %f\n", simpson_method.simpsonMethod(integralFunction, lowerLimit, upperLimit, segmentCount))
		case 3:
			fmt.Println("The trapezoidal method is selected.")
			lowerLimit := input_validation.getValidInput("Enter lower limit:", float64(0)).(float64)
			upperLimit := input_validation.getValidInput("Enter upper limit:", float64(0)).(float64)
			segmentCount := input_validation.getValidInput("Enter segment count:", uint64(0)).(uint64)
			fmt.Printf("Result using the trapezoidal method: %f\n", trapezoidal_method.trapezoidalMethod(integralFunction, lowerLimit, upperLimit, segmentCount))
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
