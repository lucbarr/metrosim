package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	womenRate  = 0.58
	pinkChance = 0.96 // chance of women chosing for pink wagon
)

type Simulator struct {

	// parameters
	EntryRate      [3]int  // rate of entry, in minutes
	WagonInterval  float64 // 1/lambda, which is expected value in minutes
	PinkRatio      float64 // ratio of pink wagons
	DroppingChance float64 // chance of dropping at a station
	TotalCapacity  int

	// state variables
	NumberOfWomen  [3]int // women in each station
	NumberOfMen    [3]int // men in each station
	CurrentStation int

	WomenInNormalWagon int
	MenInNormalWagon   int
	WomenInPinkWagon   int
}

func (s *Simulator) String() string {
	var sb strings.Builder
	for i := 0; i < 3; i++ {
		sb.WriteString(fmt.Sprintf("station %v : women %v men %v\n", i, s.NumberOfMen[i], s.NumberOfWomen[i]))
	}

	sb.WriteString(fmt.Sprintf("normal women %v men %v pink %v", s.WomenInNormalWagon, s.MenInNormalWagon, s.WomenInPinkWagon))

	return sb.String()
}

// Run updates the simulator state by n steps
func (s *Simulator) Run(maxSteps int, maxTime float64) {
	var time float64
	for i := 0; i < maxSteps; i++ {
		fmt.Println(s)
		time += s.Step()
		fmt.Println("time: ", time)
		if time > maxTime {
			break
		}
	}
}

// Step performs a state update
func (s *Simulator) Step() float64 {
	time := s.nextWagonWait()

	// update entry
	for i := 0; i < 3; i++ {
		entry := int(float64(s.EntryRate[i]) * time)

		var women, men int
		for j := 0; j < entry; j++ {
			if rand.Float64() > womenRate {
				women++
			} else {
				men++
			}
		}

		s.NumberOfMen[i] += men
		s.NumberOfWomen[i] += women
	}

	// drop from wagon
	womenNormalDrops := 0
	for i := 0; i < s.WomenInNormalWagon; i++ {
		if rand.Float64() < s.DroppingChance {
			womenNormalDrops++
		}
	}

	menNormalDrops := 0
	for i := 0; i < s.MenInNormalWagon; i++ {
		if rand.Float64() < s.DroppingChance {
			menNormalDrops++
		}
	}

	pinkDrops := 0
	for i := 0; i < s.WomenInPinkWagon; i++ {
		if rand.Float64() < s.DroppingChance {
			pinkDrops++
		}
	}

	s.WomenInNormalWagon -= womenNormalDrops
	s.MenInNormalWagon -= menNormalDrops
	s.WomenInPinkWagon -= pinkDrops

	// get into wagon
	var pinkIn, normalWomenIn int
	for i := 0; i < s.NumberOfWomen[s.CurrentStation]; i++ {
		if rand.Float64() < pinkChance {
			pinkIn++
		} else {
			normalWomenIn++
		}
	}

	s.WomenInNormalWagon += normalWomenIn
	s.MenInNormalWagon += s.NumberOfMen[s.CurrentStation]
	s.WomenInPinkWagon += pinkIn

	s.NumberOfMen[s.CurrentStation] = 0
	s.NumberOfWomen[s.CurrentStation] = 0

	s.CurrentStation = (s.CurrentStation + 1) % 3

	return time
}

func (s *Simulator) nextWagonWait() float64 {
	return rand.ExpFloat64() * s.WagonInterval
}
