package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var filename string
	var slice []Name

	fmt.Printf("Enter the filename: ")
	fmt.Scan(&filename)

	file, e := os.ReadFile(filename)
	if e != nil {
		log.Fatal(e)
	}

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		var newName Name
		splitted := strings.Split(line, " ")
		newName.fname = splitted[0]
		newName.lname = splitted[1]
		slice = append(slice, newName)
	}

	for _, item := range slice {
		fmt.Printf("%s %s\n", item.fname, item.lname)
	}

}