package main

import (
	"fmt"
)

/**
 * a race condition is when one or more operation are expected to be executed in a specific order
 * but the programm havan't be written in a way that guarented to be excuted in this order.
 *
 * go run -race raceCond.go
 * the value is  1
 * ==================
 * WARNING: DATA RACE
 * Write at 0x00c0000b2010 by goroutine 7:
 *  main.main.func1()
 *   /path/raceCond.go:15 +0x4e
 *  Previous read at 0x00c0000b2010 by main goroutine:
 *   main.main()
 *   /path/raceCond.go:18 +0x99
 *  Goroutine 7 (running) created at:
 *   main.main()
 *   /path/raceCond.go:14 +0x8b
 * ==================
 * the value is  2
 * Found 1 data race(s)
 * exit status 66
 */
func main() {

	data := 1
	go func() {
		data++
	}()

	fmt.Println("the value is ", data)
	go func() {
		data = data + 2
	}()

	fmt.Println("the value is ", data)
}
