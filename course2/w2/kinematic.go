package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acc, iVelo, iDisp float64) func(time float64) float64 {
	return func(time float64) float64 {
		return 0.5*acc*math.Pow(time, 2) + iVelo*time + iDisp
	}
}

func main() {
	var acc, iVelo, iDisp, time float64

	fmt.Print("Please enter acceleration: ")
	fmt.Scanln(&acc)

	fmt.Print("Please enter initial velocity: ")
	fmt.Scanln(&iVelo)

	fmt.Print("Please enter initial displacement: ")
	fmt.Scanln(&iDisp)

	fmt.Print("Please enter time: ")
	fmt.Scanln(&time)

	displacementFn := GenDisplaceFn(acc, iVelo, iDisp)
	fmt.Println("Displacement:", displacementFn(time))
}
