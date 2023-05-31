package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }

func (a Cow) Eat()   { fmt.Println("The cow eats grass") }
func (a Cow) Move()  { fmt.Println("The cow walk") }
func (a Cow) Speak() { fmt.Println("The cow moo") }

type Bird struct{ name string }

func (a Bird) Eat()   { fmt.Println("The bird eats worms") }
func (a Bird) Move()  { fmt.Println("The bird fly") }
func (a Bird) Speak() { fmt.Println("The bird peep") }

type Snake struct{ name string }

func (a Snake) Eat()   { fmt.Println("The snake eats mice") }
func (a Snake) Move()  { fmt.Println("The snake slither") }
func (a Snake) Speak() { fmt.Println("The snake hsss") }

func main() {
	animals := make(map[string]Animal)
	var a Animal
	var command string
	var name string
	var type_or_action string

	for {
		fmt.Printf("> ")
		fmt.Scan(&command)
		fmt.Scan(&name)
		fmt.Scan(&type_or_action)

		if command == "newanimal" {
			if type_or_action == "cow" {
				a = Cow{name}
			} else if type_or_action == "bird" {
				a = Bird{name}
			} else if type_or_action == "snake" {
				a = Snake{name}
			} else {
				fmt.Println("Invalid animal!")
				break
			}

			animals[name] = a
			fmt.Println("Created it!")

		} else if command == "query" {
			animal, exists := animals[name]
			if exists == true {
				if type_or_action == "eat" {
					animal.Eat()
				} else if type_or_action == "move" {
					animal.Move()
				} else if type_or_action == "speak" {
					animal.Speak()
				} else {
					fmt.Println("Invalid action!")
					break
				}
			}

		} else {
			fmt.Println("Invalid command!")
			break
		}
	}

}
