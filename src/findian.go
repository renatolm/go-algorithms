package main
import  (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Printf("Enter a string: ")
	fmt.Scan(&input)

	lower := strings.ToLower(input)

	if strings.HasPrefix(lower,"i") && strings.HasSuffix(lower,"n") && strings.Contains(lower, "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
	
}