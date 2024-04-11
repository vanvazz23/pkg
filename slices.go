package pkg

import (
	"fmt"
	"os"
	"strings"
)

var emailWriterChan chan string // Channel to send extracted emails to the writer goroutine

// UniqueStrings returns a new slice with unique strings from the given slice.
func UniqueStrings(input []string) []string {
	seen := make(map[string]struct{}) // Using an empty struct{} to minimize memory usage
	var result []string

	for _, value := range input {
		if _, exists := seen[value]; !exists {
			seen[value] = struct{}{}
			result = append(result, value)

			// Save the email to the file
			if emailWriterChan != nil {
				emailWriterChan <- value
			}
		}
	}

	return result
}

// StartEmailWriter starts the goroutine to write emails to a file.
func StartEmailWriter(filePath string) {
	emailWriterChan = make(chan string)

	go func() {
		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		for email := range emailWriterChan {
			_, err := file.WriteString(email + "\n")
			if err != nil {
				panic(err)
			}
		}
	}()
}
