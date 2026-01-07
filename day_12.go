package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day_12(print bool) {

	// this code will not provide a correct answer in the general case
	// the input is trivial to solve, so only need to do the basic first pass
	// to yield the correct number of fillable spaces.
	// intention was to produce a bitmask-type iterative shape placer.

	test_presents, test_spaces := load_day_12_data("day_12_test.csv")
	real_presents, real_spaces := load_day_12_data("day_12.csv")

	// Part 1
	test_out := process_part1_day12(test_presents, test_spaces)
	out := process_part1_day12(real_presents, real_spaces)

	if print {
		fmt.Println()
		fmt.Println()
		fmt.Println("** Day 12 **")
		fmt.Println("// Part 1 //")
		fmt.Println("Part 1 test output: ", test_out)
		fmt.Println("Part 1 real output: ", out)
	}

}

func get_valid_placements(present_no int, presents map[int]present_shape) [][]int {
	// returns a slice of slices of ints.
	// index in the top level slice refers to the locations
	// around a shape of value int.
	// 0  1  2  3  4
	// 5  6  7  8  9
	// 10 11 12 13 14
	// 15 16 17 18 19
	// 20 21 22 23 24
	// index 12 is where the current present is placed.

	out := [][]int{}

	shape := `110
	 		  011
	 		  110`

	fmt.Println(shape)

	return out

}

func process_part1_day12(presents map[int]present_shape, spaces []map[present_space][]int) int {

	// quick test to remove obvious fail and pass cases
	total_passed := 0
	for _, space := range spaces {
		if result := assess_setup(space); result != -1 {
			total_passed += result
		} else {
			fmt.Println("couldn't determine")
		}
	}

	// Plan

	// Helper function which minimises the sum of the dimensions of two shapes when added together
	// -> so that we can iteratively add new shapes to the space, increasing the overall dimension
	// -> of the space as slowly as possible

	// Use this helper function to work out which combination of two shapes has the minimum total 
	// -> dimension. So we can then add pairs of these shapes at the same time

	// Helper function which can add two shapes together / check whether the location of two shapes
	// overlap.


	return total_passed
}

func assess_setup(space map[present_space][]int) int {
	// return 1 if the space is fillable because there are sufficient 3x3 region to fit all within
	// return 0 if the space is not fillable because the total area of the shapes is greater than the space
	// return -1 if cannot be determined using these simple rules
	
	box_count := 0
	box_area := 0
	space_area := 0
	space_3x3_count := 0
	for present_space, boxes := range space {
		for _, box := range boxes {
			box_count += box
			// use 7 as an approximate average of the box area
			box_area += 9
		}
		space_area += present_space.width * present_space.length
		space_3x3_count += (present_space.width / 3) * (present_space.length / 3)
	}

	if box_count > space_3x3_count {
		return 0
	} else {
		return 1
	}
	// if box_area < space_area {
	// 	return 1
	// }
	// if box_area > space_area {
	// 	return 0
	// }
	// if box_area < space_area {
	// 	return 1
	// }
	// return -1
}

type present_space struct {
	width, length int
}

type present_shape struct {
	r1, r2, r3 []int
}

func load_day_12_data(filename string) (map[int]present_shape, []map[present_space][]int) {
		f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	present_shapes := make(map[int]present_shape)
	present_spaces := []map[present_space][]int{}

	var scanning_presents bool
	var current_present_scanning, current_present_row int

	for scanner.Scan() {
		
		row := scanner.Text()

		if row == "" {
			scanning_presents = false
			continue
		}
		
		if row[len(row)-1] == ':' {
			scanning_presents = true
			current_present_row = 0
			id, idErr := strconv.Atoi(string(row[0]))
			if idErr != nil {
				log.Fatal(idErr)
			}
			current_present_scanning = id
			continue
		}
		if scanning_presents {
			new_row := []int{}
			for _, char := range row {
				if char == '#' {
					new_row = append(new_row, 1)
				} else {
					new_row = append(new_row, 0)
				}
			}
			if current_present_row == 0 {
				present_shapes[current_present_scanning] = present_shape{
					r1: new_row,
				}
			} else if current_present_row == 1 {
				x := present_shapes[current_present_scanning]
				x.r2 = new_row
				present_shapes[current_present_scanning] = x
			} else if current_present_row == 2 {
				x := present_shapes[current_present_scanning]
				x.r3 = new_row
				present_shapes[current_present_scanning] = x
			} else {
				log.Fatal("something went wrong")
			}
			current_present_row ++
			continue
		}
		parts := strings.Fields(row)

		dims := strings.Split(parts[0][:len(parts[0])-1],"x")
		w, wErr := strconv.Atoi(dims[0])
		if wErr != nil {
			log.Fatal(wErr)
		}
		l, lErr := strconv.Atoi(dims[1])
		if lErr != nil {
			log.Fatal(lErr)
		}
		
		sp := present_space{width: w, length: l}
		
		counts := []int{}
		for _, count := range parts[1:] {
			c_i, cErr := strconv.Atoi(count)
			if cErr != nil {
				log.Fatal(cErr)
			}
			counts = append(counts, c_i)
		}
		out_map := make(map[present_space][]int)
		out_map[sp] = counts

		present_spaces = append(present_spaces, out_map)
		
	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return present_shapes, present_spaces
}