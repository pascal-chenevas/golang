package main

import (
	"fmt"
)

//
// Datatype for all SLICES []int
//
type sliceInt []int

//
// FUNCTION SWAP: to exchange elements in a slice
// INPUT: slice type sliceInt
// OUTPUT:
//
func swap(sli sliceInt) {
	sli[0], sli[1] = sli[1], sli[0]
}

//
// FUNCTION CICLOINTER: bubblesort verification cycle
// INPUT: slice type sliceInt
// OUTPUT: bool to know if there was an exchange
//
func cicloInter(sli sliceInt) bool {
	n := len(sli)
	inter := false
	for i := 1; i < n; i++ {
		if sli[i-1] > sli[i] {
			swap(sli[i-1 : i+1])
			inter = true
		}
	}
	return inter
}

//
// FUNCTION BUBBLESORT: bubblesort main cycle
// INPUT: slice type sliceInt
// OUTPUT:
//
func bubbleSort(sli sliceInt) {
	n := len(sli)
	inter := cicloInter(sli)
	for inter {
		inter = cicloInter(sli[0 : n-1])
	}
}

//
// FUNCTION CARGASLIDE: request 10 int values, return them on a slice sliceInt
// INPUT:
// OUTPUT: slice type sliceInt
//
func cargaSlice() sliceInt {
	var sli sliceInt
	var num int
	for i := 1; i <= 10; i++ {
		fmt.Printf("Enter integer or %v/10: ", i)
		fmt.Scan(&num)
		sli = append(sli, num)
	}
	return sli
}

func main() {
	sli := cargaSlice()
	fmt.Print("Original >> ")
	fmt.Println(sli)
	bubbleSort(sli)
	fmt.Print("Ordered  >> ")
	fmt.Println(sli)
}
