package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/akhilmk/GoCodes/LongestPrefixFinder/config"
)

// Program's entry point.
func main() {
	fmt.Println("--- Main Start ---")

	// Loading configurations.
	config.LoadAppConfig()

	done := make(chan string)

	go findLongestPrefix(config.GetInputValue(), done)

	longestPrefix := <-done

	if len(longestPrefix) > 0 {
		fmt.Println("Input Value: " + config.GetInputValue())
		fmt.Println("Longest Prefix: " + longestPrefix)
	} else {
		fmt.Println("No Matching Prefix Found")
	}

	fmt.Println("--- Main End ---")
}

// Function to find langest prefix of given string from the prefix file.
func findLongestPrefix(inputString string, done chan<- string) {
	startTime := time.Now()

	// Read file
	file, err := os.Open(config.GetPrefixFilePath())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	longestPrefix := ""
	scanner := bufio.NewScanner(file)

	// Linearly scanning each line.
	// If the file is large in size - do the code inside Scan() block with multiple goroutines as different batch.
	// Ex: Devide file in 4 parts and let 4 goroutine do the processing, out of 4 output make the longest as final output.
	for scanner.Scan() {

		// Check whether the current prefix found in the beginning of the given string (0th index).
		if strings.Index(inputString, scanner.Text()) == 0 {
			// Updating the longest prefix value if current prefix is longer than last found prefix.
			// If two prefixes have the same length, the recent one will be considered.
			if len(scanner.Text()) > len(longestPrefix) {
				longestPrefix = scanner.Text()
			}
		}
	}

	endTime := time.Now()
	// Performance loggging
	fmt.Println("Time Taken: " + endTime.Sub(startTime).String())

	// Sending longest prefix back.
	done <- longestPrefix
}
