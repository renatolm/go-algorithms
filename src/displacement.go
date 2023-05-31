package main

import (
	"fmt"
)

func GenDisplaceFn(a float64, v0 float64, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * t * t) + (v0 * t) + s0
	}
}

func main() {
	var a float64
	var v0 float64
	var s0 float64
	var t float64

	fmt.Printf("Enter a value for acceleration: ")
	fmt.Scan(&a)

	fmt.Printf("Enter a value for initial velocity: ")
	fmt.Scan(&v0)

	fmt.Printf("Enter a value for initial displacement: ")
	fmt.Scan(&s0)

	fmt.Printf("Enter a value for time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Println(fn(t))

}
