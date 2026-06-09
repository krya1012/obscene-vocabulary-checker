package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filename string
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	taboo := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			taboo[strings.ToLower(word)] = true
		}
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Bye!")
			return
		}

		if taboo[strings.ToLower(input)] {
			fmt.Println(strings.Repeat("*", len(input)))
		} else {
			fmt.Println(input)
		}
	}
}
