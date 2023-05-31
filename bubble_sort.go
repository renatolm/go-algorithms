package main
import (
	"fmt"
)

func Swap(arr []int, i int) {
	aux := arr[i+1]
	arr[i+1] = arr[i]
	arr[i] = aux
}

func BubbleSort(arr []int) {
	for j:=len(arr)-1; j>=0; j-- {
		for i:=0;i<j;i++ {
			if arr[i]>arr[i+1]{
				Swap(arr, i)
			}
		}
	}
	
}

func main() {
	
	arr := make([]int, 10)
	

	fmt.Println("Enter the array of numbers to be sorted: ")
	for i:=0; i<10; i++{
		fmt.Scanf("%d", &arr[i])
	}

	BubbleSort(arr)

	fmt.Println(arr)
}