package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var slice = make([]int, 0, 3)

	var option string

	for option != "X" {

		fmt.Printf("Enter an integer or 'X' to exit: ")
		fmt.Scan(&option)

		converted, e := strconv.Atoi(option)
    	if e == nil {
    		//option converted to int without error
        	slice = append(slice, converted)
        	sort.Ints(slice)
        	fmt.Printf("Sorted slice: %v\n", slice)
    	}

	}
}