package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Returns an int >= min, < max
func randomInt(min, max int) int {
    return min + rand.Intn(max-min)
}

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println(randomInt(1, 11)) //get an int in the 1...10 range
}
