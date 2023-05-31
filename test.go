package main

import (
    "fmt"
    "strings"
)


func main() {
    s := "oi tchau bença"
 
    splitted := strings.Split(s, " ") 

    fmt.Printf("%s", splitted[0])
}