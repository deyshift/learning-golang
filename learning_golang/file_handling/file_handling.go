package file_handling

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadDataFromFile(fileName string) (float64, error) {
	// Read the account balance from a file
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueText, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 1000, errors.New("failed to read data from file")
	}

	return valueText, nil
}

func WriteDataToFile(value float64, fileName string) {
	// Convert the balance to a string
	valueStr := fmt.Sprintf("%.2f", value)

	// Write the account balance to a file
	err := os.WriteFile(fileName, []byte(valueStr), 0644)
	if err != nil {
		fmt.Println("Error writing balance to file:", err)
		return
	}

	fmt.Println("Balance successfully written to file.")
}
