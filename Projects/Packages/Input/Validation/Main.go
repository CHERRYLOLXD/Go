package input_validation

import (
	"bufio"
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
		number, error := strconv.ParseUint(input, 10, 8)
		if error != nil {
			fmt.Println("Invalid input. Please enter a valid |uint8|.")
			return getValidInput(prompt, typeValue)
		}
		return uint8(number)
	case uint64:
		number, error := strconv.ParseUint(input, 10, 64)
		if error != nil {
			fmt.Println("Invalid input. Please enter a valid |uint64|.")
			return getValidInput(prompt, typeValue)
		}
		return uint64(number)
	case float64:
		number, error := strconv.ParseFloat(input, 64)
		if error != nil {
			fmt.Println("Invalid input. Please enter a valid |float64|.")
			return getValidInput(prompt, typeValue)
		}
		return float64(number)
	default:
		fmt.Println("Unsupported data type")
		return nil
	}
}
