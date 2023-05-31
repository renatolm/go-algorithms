package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	m := make(map[string]string)

	var name string
	var address string

	fmt.Printf("Enter a name: ")
	fmt.Scan(&name)

	fmt.Printf("Enter an address: ")
	fmt.Scan(&address)

	m["name"] = name
	m["address"] = address

	result, e := json.Marshal(m)
	if e == nil {
		fmt.Printf("%s\n",string(result))
	}	

}