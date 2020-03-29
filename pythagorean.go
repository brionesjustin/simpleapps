package main

import (
	"fmt"
)

type Triplet [3]int

func Range(min, max int) []Triplet {
	tripslice := []Triplet{}
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			for c := b + 1; c <= max; c++ {
				if a*a + b*b == c*c {
					tripslice = append(tripslice, Triplet{a, b, c})
				}
			}
		}
	}
	return tripslice 
}


func sumtrip(p int) []Triplet {
	newslice := []Triplet{}

	rr := Range(1, p-3)

	for _, t := range rr {
		if t[0] + t[1] + t[2] == p {
			newslice = append(newslice, t)
		}
	
	}
	return newslice


}

func main() {
	fmt.Println(sumtrip(1000))
}
