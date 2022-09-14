package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(entry string, operator string) int {
	cleanEntry := strings.Split(entry, operator)
	operator1 := parser(cleanEntry[0])
	operator2 := parser(cleanEntry[1])

	switch operator {
	case "+":
		return operator1 + operator2
	case "-":
		return operator1 - operator2
	case "*":
		return operator1 * operator2
	case "/":
		return operator1 / operator2
	default:
		return 0
	}
}

func parser(entry string) int {
	operator, _ := strconv.Atoi(entry)

	return operator
}

func readEntry() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func main() {
	entry := readEntry()
	operator := readEntry()

	c := calc{}

	fmt.Println(c.operate(entry, operator))
}
