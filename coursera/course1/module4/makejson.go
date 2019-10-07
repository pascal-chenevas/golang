package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	person := make(map[string]string)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name:")
	name, _ := reader.ReadString('\n')
	name = strings.Trim(name, "\n")

	fmt.Print("Enter your address:")
	address, _ := reader.ReadString('\n')
	address = strings.Trim(address, "\n")

	person["name"] = name
	person["address"] = address

	json, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error can not indent the given map")
	} else {
		fmt.Println(string(json))
	}

}
