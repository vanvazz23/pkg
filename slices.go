package pkg

import (
	"fmt"
	"os"
	"strings"
)

var emailWriterChan chan string

func StartEmailWriter(path string) {
	emailWriterChan = make(chan string)

	go func() {
		file, err := os.Create(path)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		for email := range emailWriterChan {
			if _, err := file.WriteString(email + "\n"); err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
	}()
}

// IsEqualSlice checks if two slices of strings are equal.
func IsEqualSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// UniqueStrings returns a new slice with unique strings from the given slice.
func UniqueStrings(input []string) []string {
	seen := make(map[string]struct{}) // Using an empty struct{} to minimize memory usage
	var result []string

	for _, value := range input {
		if _, exists := seen[value]; !exists {
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}

// FilterOutCommonExtensions filters out common file extensions from a slice of strings.
func FilterOutCommonExtensions(input []string) []string {
	exts := []string{".png", ".jpg", ".jpeg", ".gif", ".css", ".js", ".ico", ".svg", ".webp", ".pdf", ".zip", ".rar", ".tar", ".gz", ".7z", ".mp3", ".mp4", ".avi", ".mkv", ".mov", ".wmv", ".flv", ".m4v", ".webm", ".ogg", ".flac", ".wav", ".aac", ".wma", ".m4a", ".opus", ".mid", ".midi", ".mpg", ".mpeg", ".m4v", ".wmv", ".flv", ".m4v", ".webm", ".ogg", ".flac", ".wav", ".aac", ".wma", ".m4a", ".opus", ".mid", ".midi", ".mpg", ".mpeg"}

	filtered := []string{}
	for _, file := range input {
		hasCommonExtension := false
		for _, ext := range exts {
			if strings.HasSuffix(strings.ToLower(file), ext) {
				hasCommonExtension = true
				break
			}
		}
		if !hasCommonExtension {
			filtered = append(filtered, file)
		}
	}
	return filtered
}
