package main

import "fmt"
import "strconv"
import "strings"
import "sort"

func main() {
	var s string
	sli := make([]int, 3)

	for {
		fmt.Printf("Please enter either an integer the charager X: ")
		_, err := fmt.Scan(&s)
		if err != nil {
			continue
		}
		if strings.Compare(strings.ToLower(strings.TrimSpace(s)), "x") == 0 {
			break
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		} else {
			sli = append(sli, n)
			sort.Ints(sli)
		}
	}

	fmt.Println(sli)

}
