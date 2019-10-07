package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	input := ""
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your text: ")
	for scanner.Scan() {
		input = scanner.Text()
		fmt.Print("")
		re := regexp.MustCompile(`^[iI].*a+.*[nN]$`)
		if re.MatchString(input) {
			fmt.Printf("Found!")
		} else {
			fmt.Printf("Not Found!")
		}
		fmt.Println("")
		fmt.Print("Enter your text: ")
	}
}
