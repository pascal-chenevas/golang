package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Persons struct {
	Persons []Person `json:"persons"`
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Steps struct {
	Steps []Step `json:"step"`
}

func main() {

	jsonFile, err := os.Open("inline.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened inline.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var persons Persons

	json.Unmarshal(byteValue, &persons)

	for i := 0; i < len(persons.Persons); i++ {
		fmt.Println("Name: " + persons.Persons[i].Lastname + ", " + persons.Persons[i])
		fmt.Println("Age: " + strconv.Itoa(persons.Persons[i].Age))
	}
}
type Person struct {
	Id        string `json:"_id`
	Index     int    `json:"index"`
	Guid      string `json:"guid"`
	IsActive  bool   `json:"isActive"`
	Balance   string `json:"balance"`
	Picture   string `json:"picture"`
	Age       int    `json:"age"`
	EyeColor  string `json:"eyeColor"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {

	jsonFile, err := os.Open("person.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened person.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var persons Persons

	json.Unmarshal(byteValue, &persons)

	for i := 0; i < len(persons.Persons); i++ {
		fmt.Println("Name: " + persons.Persons[i].Lastname + ", " + persons.Persons[i])
		fmt.Println("Age: " + strconv.Itoa(persons.Persons[i].Age))
	}
}
