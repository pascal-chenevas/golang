package main

import (
	"fmt"
	"time"
)

func worker(queue chan int, worknumber int, done, ks chan bool) {
	for true {
		select {
		case k := <-queue:
			fmt.Println("doing work!", k, "worknumber", worknumber)
			done <- true
		case <-ks:
			fmt.Println("worker halted, number", worknumber)
			return
		}
	}
}

func main() {

	//channel for terminating the workers
	killsignal := make(chan bool)

	//queue of jobs
	q := make(chan int)
	// done channel takes the result of the job
	done := make(chan bool)

	numberOfWorkers := 2
	for i := 0; i < numberOfWorkers; i++ {
		go worker(q, i, done, killsignal)
	}

	numberOfJobs := 17
	for j := 0; j < numberOfJobs; j++ {
		go func(j int) {
			q <- j
		}(j)
	}

	// a deadlock occurs if c >= numberOfJobs
	for c := 0; c < numberOfJobs; c++ {
		<-done
	}

	fmt.Println("finished")

	// cleaning workers
	close(killsignal)
	time.Sleep(2 * time.Second)
}
