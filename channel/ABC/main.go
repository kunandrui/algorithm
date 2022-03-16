package main

import (
	"fmt"
	"sync"
)

var count = 5

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)

	chanA := make(chan struct{}, 1)
	chanB := make(chan struct{}, 1)
	chanC := make(chan struct{}, 1)

	chanA <- struct{}{}

	go printA(wg, chanA, chanB)
	go printB(wg, chanB, chanC)
	go printC(wg, chanC, chanA)

	wg.Wait()
}

func printA(wg *sync.WaitGroup, chanA chan struct{}, chanB chan struct{}) {
	defer wg.Done()

	for i := 0; i < count; i++ {
		<-chanA
		fmt.Printf("%d: A\n", i)
		chanB <- struct{}{}
	}
}

func printB(wg *sync.WaitGroup, chanB chan struct{}, chanC chan struct{}) {
	defer wg.Done()

	for i := 0; i < count; i++ {
		<-chanB
		fmt.Printf("%d: B\n", i)
		chanC <- struct{}{}
	}
}

func printC(wg *sync.WaitGroup, chanC chan struct{}, chanA chan struct{}) {
	defer wg.Done()

	for i := 0; i < count; i++ {
		<-chanC
		fmt.Printf("%d: C\n", i)
		chanA <- struct{}{}
	}
}
