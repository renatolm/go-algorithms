package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(slice []int) int {
	if len(slice) == 0 {
		return 0
	} else {
		return slice[0] + sum(slice[1:])
	}
}

func main() {
	var slice []int
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the list of integers to sum: ")
	scanner.Scan()

	input = scanner.Text()

	// words := strings.Fields(input)
	for _, num := range strings.Fields(input) {
		j, _ := strconv.Atoi(num)
		slice = append(slice, j)
	}

	fmt.Println(sum(slice))
}
