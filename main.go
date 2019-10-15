package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	sim := &Simulator{
		EntryRate:      [3]int{100, 100, 100},
		WagonInterval:  4.0,
		PinkRatio:      0.3,
		DroppingChance: 0.5,
		TotalCapacity:  10000,
	}

	sim.Run(100)
}
