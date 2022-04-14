/*
Write a program which allows the user to create a set of animals and to get information about those animals.
Each animal has a name and can be either a cow, bird, or snake.
With each command, the user can either create a new animal of one of the three types,
or the user can request information about an animal that he/she has already created.
Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type,
but the types of animals are restricted to either cow, bird, or snake.
The following table contains the three types of animals and their associated data.

+--------+------------+-------------------+--------------+
| Animal | Food eaten | Locomotion method | Spoken sound |
+--------+------------+-------------------+--------------+
| cow    | grass      | walk              | moo          |
| bird   | worms      | fly               | peep         |
| snake  | mice       | slither           | hsss         |
+--------+------------+-------------------+--------------+

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program should accept one command at a time from the user, print out a response,
and print out a new prompt on a new line. Your program should continue in this loop forever.
Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings.
The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal.
The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”.
The second string is the name of the animal.
The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion,
and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake.
For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake
all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type.
Your program should call the appropriate method when the user issues a query command.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

type Bird struct{}

type Snake struct{}

func (animal Cow) Eat() {
	fmt.Println("grass")
}

func (animal Cow) Move() {
	fmt.Println("walk")
}

func (animal Cow) Speak() {
	fmt.Println("moo")
}

func (animal Bird) Eat() {
	fmt.Println("worms")
}

func (animal Bird) Move() {
	fmt.Println("fly")
}

func (animal Bird) Speak() {
	fmt.Println("peep")
}

func (animal Snake) Eat() {
	fmt.Println("mice")
}

func (animal Snake) Move() {
	fmt.Println("slither")
}

func (animal Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := map[string]Animal{}
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("> ")
		scanner.Scan()
		fields := strings.Fields(scanner.Text())
		if len(fields) != 3 {
			fmt.Println("Incorrect number of arguments")
			continue
		}
		name := fields[1]
		switch command := fields[0]; command {
		case "newanimal":
			if _, ok := animals[name]; ok {
				fmt.Println("Animal with such name already exists")
				continue
			}
			switch animalType := fields[2]; animalType {
			case "cow":
				animals[name] = Cow{}
			case "bird":
				animals[name] = Bird{}
			case "snake":
				animals[name] = Snake{}
			default:
				fmt.Println("Unsupported animal type")
			}
			fmt.Println("Created it!")
		case "query":
			switch action := fields[2]; action {
			case "eat":
				if _, ok := animals[name]; ok {
					animals[name].Eat()
				}
			case "move":
				if _, ok := animals[name]; ok {
					animals[name].Move()
				}
			case "speak":
				if _, ok := animals[name]; ok {
					animals[name].Speak()
				}
			default:
				fmt.Println("Unsupported action")
			}
		default:
			fmt.Println("Unsupported command")
		}
	}
}
