package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var filename string
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
