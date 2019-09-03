package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var number float64
	result := ""

	for {
		fmt.Printf("Number: ")
		fmt.Scan(&number)
		trunc := int(math.Round(number))
		result = strconv.Itoa(trunc)
		fmt.Printf("Truncated: %s", result)
		fmt.Println("")
	}
}
