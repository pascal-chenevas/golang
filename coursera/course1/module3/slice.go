package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	s "strings"
)

func main() {

	input := ""
	msg := "Enter a number: "
	sli := make([]int, 0, 3)
	scanner := bufio.NewScanner(os.Stdin)

	separator := "---------"

	fmt.Println("")
	fmt.Println(separator)
	fmt.Print(msg)

	i := 0
	for scanner.Scan() {
		input = scanner.Text()

		if s.ToLower(input) == "x" {
			fmt.Println("Exited")
			os.Exit(0)
		}
		number, err := strconv.Atoi(input)

		if err == nil {
			sli = append(sli, number)
			sort.Ints(sli)
			fmt.Printf("slice : %v", sli)
			fmt.Println("")
			fmt.Println(separator)
			fmt.Print(msg)
			i++
		} else {
			fmt.Printf("%s is not an int !", input)
			fmt.Println("")
			fmt.Print(msg)
		}

	}
}
