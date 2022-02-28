package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var waitSum float64
	for i := 0; i < 10000; i++ {
		waitSum += wait(5, 10)
	}
	fmt.Println(waitSum / 10000)
}

func wait(busATime float64, busBTime float64) float64 {
	waitA := rand.Float64() * busATime
	waitB := rand.Float64() * busBTime

	//fmt.Println(waitA)
	//fmt.Println(waitB)

	if waitA < waitB {
		return waitA
	} else {
		return waitB
	}
}
