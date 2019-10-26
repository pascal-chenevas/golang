package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func StringSliceToStringInt(sli []string) []int {

	sliInt := make([]int, len(sli))
	for i := 0; i < len(sli); i++ {

		s, err := strconv.Atoi(sli[i])

		if err == nil {
			sliInt[i] = s
		}
	}
	return sliInt
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	numberArray := 4
	var slices [][]int
	result := []int{}

	ch1 := make(chan []int)
	ch2 := make(chan []int)
	ch3 := make(chan []int)
	ch4 := make(chan []int)

	var subArray1 []int
	var subArray2 []int
	var subArray3 []int
	var subArray4 []int

	fmt.Print("> (Enter a list of integer): ")
	input, _ := reader.ReadString('\n')
	input = strings.Trim(input, "\n")
	input = strings.ToLower(input)

	sli := StringSliceToStringInt(strings.Split(input, " "))

	//fmt.Println(sli)
	chunkSize := (len(sli) + numberArray - 1) / numberArray

	for i := 0; i < len(sli); i += chunkSize {
		end := i + chunkSize

		if end > len(sli) {
			end = len(sli)
		}

		slices = append(slices, sli[i:end])
	}

	if len(slices) < 4 {
		fmt.Println("Pleaser enter more number because with your list I can make less subarray than 4")
		os.Exit(0)
	}
	//fmt.Println(slices)
	go func(sli []int) {
		sort.Sort(sort.IntSlice(sli))
		fmt.Println("sorted subarray #1", sli)
		ch1 <- sli
	}(slices[0])

	go func(sli []int) {
		sort.Sort(sort.IntSlice(sli))
		fmt.Println("sorted subarray #2", sli)
		ch2 <- sli
	}(slices[1])

	go func(sli []int) {
		sort.Sort(sort.IntSlice(sli))
		fmt.Println("sorted subarray #3", sli)
		ch3 <- sli
	}(slices[2])

	go func(sli []int) {
		sort.Sort(sort.IntSlice(sli))
		fmt.Println("sorted subarray #4", sli)
		ch4 <- sli
	}(slices[3])

	subArray1, subArray2, subArray3, subArray4 = <-ch1, <-ch2, <-ch3, <-ch4
	result = append(result, subArray1...)
	result = append(result, subArray2...)
	result = append(result, subArray3...)
	result = append(result, subArray4...)

	fmt.Println("====== result (sorted array) =====")
	sort.Sort(sort.IntSlice(result))
	fmt.Println(result)
}
