package main

import (
	"fmt"
	"strings"
)

func main() {
	testString := "  hello  world  "

	for _, item := range cleanInput(testString) {
		fmt.Println(item)
	}

}

func cleanInput(text string) []string {
	parts := strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
	var result []string
	for _, part := range parts {
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}
