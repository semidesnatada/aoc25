package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type safeInstrucions struct {
	Value int
	Direction bool
}

func day_1() {

	fmt.Println("first day")

	test_data := load_day_1_data("day_1_test.csv")
	real_data := load_day_1_data("day_1.csv")

	// Part 1
	zeros := process_safe_instructions(test_data)
	real_zeros := process_safe_instructions(real_data)

	fmt.Println("// Part 1 //")
	fmt.Println("Number of zeros in test data is: ", zeros)
	fmt.Println("Number of zeros in real data is: ", real_zeros)
	fmt.Println()

	// Part 2
	p2_zeros := process_safe_instructions_part_2(test_data)
	p2_real_zeros := process_safe_instructions_part_2(real_data)

	fmt.Println("// Part 2 //")
	fmt.Println("Number of zeros in test data is: ", p2_zeros)
	fmt.Println("Number of zeros in real data is: ", p2_real_zeros)
	fmt.Println()

}

func process_safe_instructions(ins []safeInstrucions) int {

	zero_count := 0
	curr := 50

	for _, instruction := range ins {

		if instruction.Direction {
			curr += (instruction.Value + 100)
		} else {
			curr -= (instruction.Value - 100)
		}
		curr = curr % 100
		if curr == 0 {
			zero_count += 1
		}

	}

	return zero_count

}

func process_safe_instructions_part_2(ins []safeInstrucions) int {

	zero_count := 0
	curr := 50

	for _, instruction := range ins {
		fmt.Println(curr, instruction.Value, instruction.Direction, zero_count)
		if (curr + instruction.Value >= 100 && instruction.Direction) || (curr - instruction.Value <= 0 && !instruction.Direction) {
			if large := curr + instruction.Value; large >= 100 && instruction.Direction {
				zero_count += large / 100
			}
			if small := curr - instruction.Value; small <= 0 && !instruction.Direction {
				if curr != 0 {
					zero_count += ((-small / 100) + 1)
				} else {
					zero_count += ((-small / 100))
				}
				
			}
		}

		// used an arbitrary large product of 100 to fix any looping.
		// probably a cleaner way to do this which handles all cases
		if instruction.Direction {
			curr += (instruction.Value + 10000)
		} else {
			curr -= (instruction.Value - 10000)
		}
		curr = curr % 100

		

	}

	return zero_count

}

func load_day_1_data(filename string) []safeInstrucions {
	f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	out := []safeInstrucions{}

	for scanner.Scan() {
		row := scanner.Text()

		direction_s := row[0]

		var dir_b bool
		if direction_s == 'L' {
			dir_b = false
		} else if direction_s == 'R' {
			dir_b = true
		} else {
			log.Fatal("error parsing direction")
		}

		val_s := row[1:]
		val_i, conv_err := strconv.Atoi(val_s)
		if conv_err != nil {
			log.Fatal(conv_err)
		}

		out = append(out, safeInstrucions{
			Value: val_i,
			Direction: dir_b,
		})

	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return out
}