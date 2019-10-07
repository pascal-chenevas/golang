package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {

	fn := func(t float64) float64 {

		return (a*math.Pow(t, 2))/2 + vo*t + so
	}

	return fn
}

func main() {

	var a, vo, so, t float64

	fmt.Print("Enter a float value (acceleration, velocity, displacement, time) : ")
	_, err := fmt.Scanf("%f %f %f %f", &a, &vo, &so, &t)

	if err != nil {
		fmt.Println(err)
	}

	fn := GenDisplaceFn(a, vo, so)
	fmt.Println(fn(t))

}
