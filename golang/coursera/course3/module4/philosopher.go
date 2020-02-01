package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

type Philosopher struct {
	number          int
	leftCs, rightCs *ChopS
}

func (p Philosopher) GetNumber() int {
	return p.number
}

func (p Philosopher) Eat(wg *sync.WaitGroup) {
	defer func() {
		cntDone()
		wg.Done()
	}()

	for i := 1; i <= 3; i++ {
		p.leftCs.Lock()
		p.rightCs.Lock()
		fmt.Println("starting to eat <", p.number, ">")
		fmt.Println("Finishing eating <", p.number, ">")
		p.leftCs.Unlock()
		p.rightCs.Unlock()
		time.Sleep(2 * time.Second)
	}

}

func NewPhilospher(number int, leftCs, rightCs *ChopS) Philosopher {
	var p Philosopher

	p.number = number
	p.leftCs = leftCs
	p.rightCs = rightCs

	return p
}

func createChops() []*ChopS {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	return CSticks
}

func createPhilosophers() []Philosopher {

	var philos []Philosopher

	CSticks := createChops()
	for i := 0; i < 5; i++ {
		philos = append(philos, NewPhilospher(i, CSticks[i], CSticks[(i+1)%5]))
	}

	return philos
}

// control number of Philosophers eating - mimicing sync.WaitGroup
var cnt int
var cntMtx sync.RWMutex

func hostCntVal() int {
	cntMtx.RLock()
	defer cntMtx.RUnlock()
	return cnt
}

func cntAdd() {
	cntMtx.Lock()
	cnt++
	cntMtx.Unlock()
}

func cntDone() {
	cntMtx.Lock()
	cnt--
	cntMtx.Unlock()
}

func main() {
	var wg sync.WaitGroup
	philos := createPhilosophers()

	for i := 0; i < 5; i++ {
	waitOnTwo:
		if hostCntVal() < 2 {
			cntAdd()
			wg.Add(1)
			go philos[i].Eat(&wg)
		} else {
			time.Sleep(7 * time.Second) // let a Philosopher finish eating
			goto waitOnTwo              // retry
		}
	}
	wg.Wait()
}
