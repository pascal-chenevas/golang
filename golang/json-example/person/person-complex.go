package main

import (
	"bufio"
	//"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	jf, err := os.Open("person-complex.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jf.Close()

	s := bufio.NewScanner(jf)

	for s.Scan() {
		fmt.Println("===========")
		fmt.Println(s.Text())
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

}
