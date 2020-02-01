package main

import (
	"fmt"
	"os"
)

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	fmt.Println("I eat ", a.food)
}

func (a Animal) Move() {
	fmt.Println("I ", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println("I ", a.noise)
}

func request(a Animal, request string) {
	if request == "eat" {
		a.Eat()
	} else if request == "move" {
		a.Move()
	} else if request == "speak" {
		a.Speak()
	} else {
		fmt.Println("request '", request, "' not supported! Allowed: eat/move/speak")
		os.Exit(0)
	}
}

func main() {

	var animalChosen Animal
	var animalName, informationRequested string

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		for i := 1; i <= 3; i++ {
			fmt.Println("request an animal (cow/bird/snake, eat/move/speak) ", i, "/3")
			fmt.Print("> ")
			_, err := fmt.Scanf("%s %s", &animalName, &informationRequested)

			if err != nil {
				fmt.Println(err)
			}

			if animalName == "cow" {
				animalChosen = cow
			} else if animalName == "bird" {
				animalChosen = bird
			} else if animalName == "snake" {
				animalChosen = snake
			} else {
				fmt.Println("animal name '", animalName, "' not allowed! Allowed: cow/bird/snake ")
				os.Exit(0)
			}
			request(animalChosen, informationRequested)
		}

	}
}
