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
	sli := make([]int, 3)
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
			if i < 3 {
				sli[i] = number
				fmt.Printf("slice : %v", sli)
				fmt.Println("")
				if i == 2 {
					sort.Ints(sli)
					fmt.Printf("sorted slice : %v", sli)
					sli = make([]int, 3)
					fmt.Println("")
					fmt.Println(separator)
					fmt.Println("")
				}
				fmt.Print(msg)
			}
			i++

		}

	}
}
