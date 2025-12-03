package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func day_3() {

	fmt.Println("third day")

	test_data := load_day_3_data("day_3_test.csv")
	real_data := load_day_3_data("day_3.csv")


	// Part 1
	test_joltage_p1 := process_part1_inputs(test_data)
	fmt.Println("Part 1 sum of test max joltages: ", test_joltage_p1)

	joltage_p1 := process_part1_inputs(real_data)
	fmt.Println("Part 1 sum of real max joltages: ", joltage_p1)

	// Part 2
	test_joltage_p2 := process_part2_inputs(test_data)
	fmt.Println("Part 2 sum of test max joltages: ", test_joltage_p2)

	joltage_p2 := process_part2_inputs(real_data)
	fmt.Println("Part 2 sum of real max joltages: ", joltage_p2)


}

func process_part1_inputs(ins [][]int) int {

	out := 0

	for _, row := range ins {
		// fmt.Println(row)
		// fmt.Println("largest joltage is: ", find_largest_joltage(row))
		out += find_largest_joltage(row)
	}

	return out
}

func process_part2_inputs(ins [][]int) int {

	out := 0

	for _, row := range ins {
		// fmt.Println(row)
		// fmt.Println("largest joltage is: ", find_largest_12_joltage(row))
		out += find_largest_12_joltage(row)
		// break
	}

	return out
}

func find_largest_12_joltage(input []int) int {

	out := 0
	current_arg := 0

	out_digs := []int{}
	// run a loop across the 12 needed digits
	for i := 0; i < 12; i++ {

		dig, arg_dig := find_max(input[current_arg:len(input)-11+i])

		out_digs = append(out_digs, dig)

		current_arg += arg_dig + 1
		out += int(math.Pow10(11-i)) * dig

	}

	// fmt.Println(out_digs)

	return out

}

func find_max(in []int) (int, int) {

	var max, argmax int

	for ind, val := range in {
		if val > max {
			max = val
			argmax = ind
		}
	}

	// fmt.Println(in)

	return max, argmax
}

func find_largest_joltage(input []int) int {

	var max, argmax int

	for ind, val := range input {
		if val > max && ind != len(input) - 1 {
			max = val
			argmax = ind
		}
	}

	var submax int//, subargmax int
	for j := argmax + 1; j < len(input); j ++ {
		if input[j] > submax {
			submax = input[j]
			// subargmax = j
		}
	}

	out := 10 * max + submax
	return out

}

func load_day_3_data(filename string) [][]int {
	f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	out := [][]int{}

	for scanner.Scan() {
		row := scanner.Text()

		outrow := []int{}

		for _, char := range row {
			// char_int, cErr := strconv.Atoi(char)
			// if cErr != nil {
			// 	log.Fatal(cErr)
			// }

			// no idea why this works for converting rune to int. look into it later.
			outrow = append(outrow, int(char-'0'))
		}

		out = append(out, outrow)

	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return out
}