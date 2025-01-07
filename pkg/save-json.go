package pkg

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveFileToJSON(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	data := Expenses
	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return fmt.Errorf("error encoding JSON: %w", err)
	}

	return nil
}
