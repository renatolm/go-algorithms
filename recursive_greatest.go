package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func greatest(max int, slice []int) int {
	if len(slice) == 0 {
		return max
	} else {
		if max < slice[0] {
			max = slice[0]
		}
		return greatest(max, slice[1:])
	}
}

func main() {
	var slice []int
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the list of integers to get the greatest: ")
	scanner.Scan()

	input = scanner.Text()

	// words := strings.Fields(input)
	for _, num := range strings.Fields(input) {
		j, _ := strconv.Atoi(num)
		slice = append(slice, j)
	}

	fmt.Println(greatest(slice[0], slice))
}
