package main

import (
	"fmt"
	"os"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func (a *Animal) InitAnimal(food, locomotion, noise string) {
	a.food = food
	a.locomotion = locomotion
	a.noise = noise
}

func main() {
	var a Animal
	// var options string
	var valid_animal bool

	animal := ""
	action := ""

	for animal != "x" && action != "x" {
		valid_animal = true

		fmt.Println("Type the animal (cow, bird, snake) and action (eat, move, speak) separated by whitespace.")
		fmt.Println("Or type 'x' to quit.")
		fmt.Scan(&animal)
		fmt.Scan(&action)

		fmt.Printf("\n")

		if animal == "cow" {
			a.InitAnimal("grass", "walk", "moo")
		} else if animal == "bird" {
			a.InitAnimal("worms", "fly", "peep")
		} else if animal == "snake" {
			a.InitAnimal("mice", "slither", "hsss")
		} else if animal == "x" {
			break
		} else {
			valid_animal = false
		}

		if action == "eat" && valid_animal == true {
			fmt.Println("The", animal, "eats", a.Eat())
		} else if action == "move" && valid_animal == true {
			fmt.Println("The", animal, a.Move())
		} else if action == "speak" && valid_animal == true {
			fmt.Println("The", animal, a.Speak())
		} else if action == "x" {
			break
		} else {
			fmt.Println("Invalid options!")
		}

		fmt.Printf("\n")

	}

	os.Exit(0)

}
