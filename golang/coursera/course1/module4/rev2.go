package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	type Name struct {
		fname string
		lname string
	}

	sli := make([]Name, 20)
	length := 20

	fmt.Println("Please input the file name")
	var fn string
	fmt.Scan(&fn)
	f, err := os.Open(fn)
	defer f.Close()

	if err != nil {
		fmt.Println("Failed to open file")
		return
	}

	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		names := strings.Fields(line)

		if i < length {
			sli[i] = Name{fname: names[0], lname: names[1]}
		} else {
			sli = append(sli, Name{fname: names[0], lname: names[1]})
		}
		i++
	}

	fmt.Println("\nAll names are:\n")
	for _, item := range sli {
		if item.fname == "" {
			break
		}
		fmt.Println(item.fname, item.lname)
	}

}
