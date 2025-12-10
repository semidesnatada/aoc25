package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start, End int
}

func day_2() {

	// fmt.Println("second day")

	test_data := load_day_2_data("day_2_test.csv")
	real_data := load_day_2_data("day_2.csv")

	fmt.Println()
	fmt.Println()
	fmt.Println("** Day 2 **")
	fmt.Println("// Part 1 //")
	// Part 1
	no_test_invalid_p1 := test_all_ranges_in_data_p1(test_data)
	fmt.Println("Part 1 test sum of invalid IDs: ", no_test_invalid_p1)

	no_invalid_p1 := test_all_ranges_in_data_p1(real_data)
	fmt.Println("Part 1 real sum of invalid IDs: ", no_invalid_p1)
	
	fmt.Println()
	fmt.Println("// Part 2 //")
	// Part 2
	no_test_invalid_p2 := test_all_ranges_in_data_p2(test_data)
	fmt.Println("Part 2 test sum of invalid IDs: ", no_test_invalid_p2)

	no_invalid_p2 := test_all_ranges_in_data_p2(real_data)
	fmt.Println("Part 2 real sum of invalid IDs: ", no_invalid_p2)


}

func test_all_ranges_in_data_p1(rs []Range) int {

	count := 0

	for _, r := range rs {

		invalids := test_double_repeat_range(r)

		for _, invalid := range invalids {
			count += invalid
		}

		// fmt.Println(invalids)

	}

	return count

}

func test_all_ranges_in_data_p2(rs []Range) int {
	count := 0
	for _, r := range rs {

		invalids := test_n_repeat_range(r)
		for _, invalid := range invalids {
			count += invalid
		}
		// fmt.Println(invalids)
	}

	return count
}

func test_n_repeat_range(r Range) []int {

	out := []int{}

	for i := r.Start; i <= r.End; i++ {
		curr := strconv.Itoa(i)

		duped := check_if_substring_is_repeated_any_times(curr)

		if duped {
			out = append(out, i)
		}
	}

	return out
}


func test_double_repeat_range(r Range) []int {

	out := []int{}

	for i := r.Start; i <= r.End; i++ {
		curr := strconv.Itoa(i)

		duped := check_if_substring_is_repeated_twice(curr)

		if duped {
			out = append(out, i)
		}
	}

	return out
}

func check_if_substring_is_repeated_any_times(s string) bool {

	if len(s) == 1 {
		return false
	}

	// found := false

	// check across all possible ways to split the string into equal parts
	for i := 2; i <= len(s); i++ {

		// check if string can be split equally
		if len(s) % i != 0 {
			continue
		}
		
		is_valid := check_n_repeats(i, s)
		// fmt.Println("string: ", s, " is valid: ", is_valid, " n repeats ", i)
		if is_valid {
			return true
		}

	}

	return false
}

func check_n_repeats(n int, s string) bool {

	// get the first part for later comparison
	base := s[ : len(s) / n]
	broken := false

	// fmt.Println("base: ", base)
	// check the base against all subsequent equal parts
	for j := 1; j < n; j++ {

		// get next equal part to compare against
		comparison := s[j*len(s)/n: (j+1)*len(s)/n]
		// fmt.Println(" subsequent: ", comparison)

		broken = checkPairIsNotEqual(base, comparison)

		if broken {
			break
		}

	}
	return !broken
}

func checkPairIsNotEqual(base, comparison string) bool {

	for i:=0; i< len(base); i++ {
		if base[i] != comparison[i] {
			return true
		}
	}
	return false

}

func check_if_substring_is_repeated_twice(s string) bool {

	if len(s) % 2 != 0 {
		return false
	}

	p1 := s[:len(s)/2]
	p2 := s[len(s)/2:]

	for i := 0; i < len(s)/2; i++ {
		if p1[i] != p2[i] {
			return false
		}
	}

	return true
}

func load_day_2_data(filename string) []Range {
	f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	out := []Range{}

	for scanner.Scan() {
		row := scanner.Text()

		ranges := strings.Split(row, ",")

		for _, rang := range ranges {
			tails := strings.Split(rang, "-")
			begin, bErr := strconv.Atoi(tails[0])
			if bErr != nil {
				log.Fatal(bErr)
			}
			end, eErr := strconv.Atoi(tails[1])
			if eErr != nil {
				log.Fatal(eErr)
			}
			out = append(out, Range{
				Start: begin,
				End: end,
			})
		}

	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return out
}