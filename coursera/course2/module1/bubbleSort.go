package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func Swap(sli []int, index int) {
	tmp := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = tmp
}

func BubbleSort(sli []int) {
	length := len(sli)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	var numbers []int
	for i := 1; i <= 10; i++ {
		fmt.Print("Enter an int ", i, "/10 : ")
		input, _ := reader.ReadString('\n')
		input = s.Trim(input, "\n")
		number, _ := strconv.Atoi(input)
		numbers = append(numbers, number)

	}
	fmt.Println("you entered: ", numbers)
	BubbleSort(numbers)
	fmt.Println("sorted list: ", numbers)
}
