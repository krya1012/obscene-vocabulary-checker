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

		tokens := strings.Fields(input)
		for i, token := range tokens {
			clean := strings.TrimRight(token, ".,!?;:'\"")
			if taboo[strings.ToLower(clean)] {
				tokens[i] = strings.Repeat("*", len(clean)) + token[len(clean):]
			}
		}
		fmt.Println(strings.Join(tokens, " "))
	}
}
