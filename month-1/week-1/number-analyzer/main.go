package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Number Analyzer CLI")
	fmt.Print("Enter a number: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	input = strings.TrimSpace(input)

	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid number input")
		return
	}

	fmt.Println("You entered:", number)
}
