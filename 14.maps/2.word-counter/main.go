package main

import (
	"fmt"
	"strings"
)

const input = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

// Implement the function count words. The function returns a map of words and their count.
// Hint: splitting by whitespace and newline can be done using: strings.Fields
func CountWords(input string) map[string]int {
	// FIXME: Implement; Run the test to verify your solution.
	wordCount := make(map[string]int)
	totalWords := strings.Fields(input)
	for _, v := range totalWords {
		_, ok := wordCount[v]
		if ok {
			wordCount[v] += 1
		} else {
			wordCount[v] = 1
		}
	}
	return wordCount
}

func main() {
	wordCountMap := CountWords(input)

	fmt.Println(wordCountMap)
}
