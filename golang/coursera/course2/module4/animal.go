package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	s "strings"
)

const CMD_NEWANIMAL = "newanimal"
const CMD_QUERY = "query"

const ANIMAL_TYPE_COW = "cow"
const ANIMAL_TYPE_BIRD = "bird"
const ANIMAL_TYPE_SNAKE = "snake"

const ANIMAL_REQUEST_EAT = "eat"
const ANIMAL_REQUEST_MOVE = "move"
const ANIMAL_REQUEST_SPEAK = "speak"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }

type Bird struct{ name string }

type Snake struct{ name string }

func (c Cow) Eat() {
	fmt.Println("I, ", c.name, ", I eat grass")
}

func (c Cow) Move() {
	fmt.Println("I, ", c.name, ", I walk")
}

func (c Cow) Speak() {
	fmt.Println("I, ", c.name, ", I moo")
}

func (b Bird) Eat() {
	fmt.Println("I, ", b.name, ", I eat worms")
}

func (b Bird) Move() {
	fmt.Println("I, ", b.name, ", I fly")
}

func (b Bird) Speak() {
	fmt.Println("I, ", b.name, ", I peep")
}

func (s Snake) Eat() {
	fmt.Println("I, ", s.name, ", I eat mice")
}

func (s Snake) Move() {
	fmt.Println("I, ", s.name, ", I slither")
}

func (s Snake) Speak() {
	fmt.Println("I, ", s.name, ", I hsss")
}

func NewAnimal(animalName, animalType string) Animal {
	animalType = s.ToLower(animalType)

	var a Animal

	if animalType == ANIMAL_TYPE_COW {
		a = Cow{animalName}
	} else if animalType == ANIMAL_TYPE_BIRD {
		a = Bird{animalName}
	} else if animalType == ANIMAL_TYPE_SNAKE {
		a = Snake{animalName}
	} else {
		fmt.Println("Animal \"", animalType, "\" can't be created because not implemented!")
		return nil
	}

	fmt.Println("new Animal \"", animalName, "\" (", animalType, ") created!")

	return a
}

func proceedRequest(a Animal, request string) {

	request = s.ToLower(request)
	if request == ANIMAL_REQUEST_EAT {
		a.Eat()
	} else if request == ANIMAL_REQUEST_MOVE {
		a.Move()
	} else if request == ANIMAL_REQUEST_SPEAK {
		a.Speak()
	} else {
		fmt.Println("Wrong request! Allowed: eat/move/speak")
	}
}

func printHelp() {
	fmt.Println("\t newanimal <animalName> <animalType=[cow|bird|snake]> : create a new animal\n\t example: newanimal foo cow")
	fmt.Println("\t query <animalName> <request=[eat|move|speak]> : display informaiton relating to an animal\n\t example: query foo eat")
}

func getAnimalByName(animalName string, animals map[string]Animal) Animal {
	var a Animal
	for k, v := range animals {
		if k == animalName {
			a = v
			break
		}
	}
	return a

}

func isCommandRunnabled(cmd string) bool {
	match, _ := regexp.MatchString(`(?i)q|h|newanimal \w+ (cow|snake|bird)|query \w+ (eat|move|speak)`, cmd)
	return match

}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var animal Animal
	animals := make(map[string]Animal)

	requests := 1

	for {

		fmt.Print("> (h for help, q to quit) ")
		input, _ := reader.ReadString('\n')
		input = s.Trim(input, "\n")
		input = s.ToLower(input)

		if isCommandRunnabled(input) {

			if input == "h" {
				printHelp()
			} else if input == "q" {
				fmt.Println("Exited!")
				os.Exit(0)
			} else {

				values := s.Split(input, " ")

				if len(values) == 3 {
					cmd := values[0]
					arg1 := values[1]
					arg2 := values[2]

					if cmd == "newanimal" && len(animals) <= 3 {
						animal := NewAnimal(arg1, arg2)
						if animal != nil {
							animals[arg1] = animal
						}
					} else if cmd == "query" && requests <= 3 {
						animal = getAnimalByName(arg1, animals)
						if animal != nil {
							proceedRequest(animal, arg2)
							requests++

						} else {
							fmt.Println("Animal \"", arg1, "\" not found!")
						}

					}

				}

			}

		}

	}
}
