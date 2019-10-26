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

	ch := make(chan []int)

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

	//fmt.Println(slices)
	for i := 0; i < len(slices); i++ {
		go func(sli []int, i int) {
			sort.Sort(sort.IntSlice(sli))
			fmt.Println("sorted subarray #:", i, sli)
			ch <- sli
		}(slices[i], i)
		result = append(result, <-ch...)
	}

	fmt.Println("====== result (sorted array) =====")
	sort.Sort(sort.IntSlice(result))
	fmt.Println(result)
}
