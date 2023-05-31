package main
import "fmt"

func main() {
	var input float32

	fmt.Printf("Enter a floating point number to be truncated: ")
	fmt.Scan(&input)

	fmt.Printf("%d\n",int(input))
}