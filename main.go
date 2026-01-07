package main

import (
	"fmt"
	"time"
)

func main () {

	// e.g.
	day_01(true)
	// day_12(true)

	all_days()

}

func all_days() {

	days := []func(bool){
		day_01, day_02, day_03,
		day_04, day_05, day_06,
		day_07, day_08, day_09,
		day_10, day_11, day_12,
	}

	t_0 := time.Now()

	for _, day := range days {
		// the bool passed into the function decides whether or not
		// the function outputs will be printed - true => print out
		day(false)
	}

	dt := time.Since(t_0)
	
	fmt.Printf("\nSolved all puzzles in %v \n\n", dt)
	

}