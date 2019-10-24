package main

import (
	"fmt"
	"sort"
)

func main() {
	sli := []int{10, 33, 5, -1, 55, 12, 13, 45, 9, 14, 534, 78, 40, -5, -6, -20, 0, 55, 60}

	var slices [4][]int
	result := []int{}

	fmt.Println(sli)
	ch := make(chan []int)

	idx := 1
	slices[0] = sli[:4]
	for i := 4; i <= len(sli); i++ {
		if i%4 == 0 && i+4 < len(sli) {
			slices[idx] = sli[i : i+4]
			idx++
		}

	}

	fmt.Println(slices)
	for i := 0; i < len(slices); i++ {
		go func(sli []int, i int) {
			sort.Sort(sort.IntSlice(sli))
			fmt.Println("slice #:", i, sli)
			ch <- sli
		}(slices[i], i)
		result = append(result, <-ch...)
	}

	fmt.Println("====== result =====")
	sort.Sort(sort.IntSlice(result))
	fmt.Println(result)
	fmt.Println(len(result))
	fmt.Println(len(sli))

}
