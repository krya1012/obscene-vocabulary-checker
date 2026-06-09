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
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if taboo[strings.ToLower(input)] {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
