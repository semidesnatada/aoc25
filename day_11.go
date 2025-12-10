package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func day_11() {


	// fmt.Println("eleventh day")

	test_machines := load_day_11_data("day_11_test.csv")
	real_machines := load_day_11_data("day_11.csv")

	fmt.Println()
	fmt.Println()
	fmt.Println("** Day 11 **")
	fmt.Println("// Part 1 //")

	// Part 1
	test_out := process_part1_day11(test_machines)
	fmt.Println("Part 1 test output: ", test_out)

	out := process_part1_day11(real_machines)
	fmt.Println("Part 1 real output: ", out)

	fmt.Println()
	fmt.Println("// Part 2 //")
	// Part 2
	// test_out := process_part1_day11(test_machines)
	// fmt.Println("Part 1 test output: ", test_out)

	// out := process_part1_day11(real_machines)
	// fmt.Println("Part 1 real output: ", out)

}

func process_part1_day11(in []int) int {

	out := 0


	return out
}


func load_day_11_data(filename string) []int {
		f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	out := []int{}

	for scanner.Scan() {
		
		// row := scanner.Text()
		

	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return out
}