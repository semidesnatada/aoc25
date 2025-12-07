package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day_5() {

	fmt.Println("fifth day")

	test_ranges, test_availables := load_day_5_data("day_5_test.csv")
	real_ranges, real_availables := load_day_5_data("day_5.csv")

	fmt.Println(test_ranges)
	fmt.Println(test_availables)

	// for i, row := range test_data {
	// 	fmt.Println(row, i)
	// }
		
	// Part 1
	test_out := process_part1_day5(test_ranges, test_availables)
	fmt.Println("Part 1 test output: ", test_out)

	out := process_part1_day5(real_ranges, real_availables)
	fmt.Println("Part 1 real output: ", out)

	// Part 2
	test_out_p2 := process_part2_day5(test_ranges)
	fmt.Println("Part 2 test output: ", test_out_p2)

	out_p2 := process_part2_day5(real_ranges)
	fmt.Println("Part 2 real output: ", out_p2)


}

func process_part2_day5(ra []ingredientRange) int {

	out := 0

	// returns a set of mutually exclusive ranges
	out_ranges := recurConsolidate(ra)
	
	for _, ran := range out_ranges {
		out += (ran.Max - ran.Min + 1)
	}

	// fmt.Println(out_ranges)

	return out

}

func recurConsolidate(ings []ingredientRange) []ingredientRange {

	for i := 0; i < len(ings); i++ {
		for j := i + 1; j < len(ings); j++ {
			// if the two ranges overlap then we need to merge them
			// and re iterate over the full list of ranges
			if can, conned := getConsolidation(ings[i], ings[j]); can {
				var new_ings []ingredientRange
				if j+1 < len(ings) {
					new_ings = slices.Concat(
						[]ingredientRange{conned},
						ings[:i], ings[i+1:j], ings[j+1:],
					)
				} else {
					new_ings = slices.Concat(
						[]ingredientRange{conned},
						ings[:i], ings[i+1:j],
					)
				}
				return recurConsolidate(new_ings)
			}
		}
	}

	return ings
}


func getConsolidation(existing, comparator ingredientRange) (bool, ingredientRange) {
	// bool indicates whether the existing input ranges overlap, true if they do

	el := existing.Min
	eh := existing.Max

	cl := comparator.Min
	ch := comparator.Max

	if eh < cl || ch < el {
		return false, comparator
	}

	// case where one range is inside the other
	if el < cl && eh > ch {
		return true, ingredientRange{Min:el, Max:eh}
	}
		if cl < el && ch > eh {
		return true, ingredientRange{Min:cl, Max:ch}
	}

	var lout, hout int

	if el < cl {
		lout = el
	} else {
		lout = cl
	}
	if eh > ch {
		hout = eh
	} else {
		hout = ch
	}

	return true, ingredientRange{
		Min: lout,
		Max: hout,
	}

}

func process_part1_day5(ra []ingredientRange, avails []int) int {

	fresh_count := 0

	for _, avail := range avails {

		for _, ing_range := range ra {

			if avail >= ing_range.Min && avail <= ing_range.Max {
				fresh_count += 1
				break
			}
		}

	}

	return fresh_count

}

type ingredientRange struct {
	Min, Max int
}

func load_day_5_data(filename string) ([]ingredientRange, []int) {
	f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	out_range := []ingredientRange{}
	out_available := []int{}

	broken := false

	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			broken = true
			continue
		}
		if !broken {
			parts := strings.Split(row, "-")
			low_i, lErr := strconv.Atoi(parts[0])
			if lErr != nil {
				log.Fatal(lErr)
			}
			hig_i, hErr := strconv.Atoi(parts[1])
			if hErr != nil {
				log.Fatal(hErr)
			}
			out_range = append(out_range, ingredientRange{
				Min: low_i,
				Max: hig_i,
			})

		} else {
			ava_int, aErr := strconv.Atoi(row)
			if aErr != nil {
				log.Fatal(aErr)
			}
			out_available = append(out_available, ava_int)

		}

	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return out_range, out_available
}