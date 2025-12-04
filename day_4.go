package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func day_4() {

	fmt.Println("fourth day")

	test_data := load_day_4_data("day_4_test.csv")
	real_data := load_day_4_data("day_4.csv")

	// for i, row := range test_data {
	// 	fmt.Println(row, i)
	// }
		
	// Part 1
	test_out, _ := process_part1_day4(test_data)
	fmt.Println("Part 1 test output: ", test_out)

	out, _ := process_part1_day4(real_data)
	fmt.Println("Part 1 real output: ", out)


	// Part 1
	test_out_p2 := process_part2_day4(test_data)
	fmt.Println("Part 2 test output: ", test_out_p2)

	out_p2 := process_part2_day4(real_data)
	fmt.Println("Part 2 real output: ", out_p2)

}

func process_part2_day4(grid []string) int {

	total_removed := 0
	var removed_in_step int

	for ;; {

		removed_in_step, grid = process_part1_day4(grid)
		total_removed += removed_in_step
		if removed_in_step == 0 {
			break
		}

	}

	return total_removed

}

func process_part1_day4(grid []string) (int, []string) {

	accesible_count := 0

	new_grid := make([]string, len(grid))
	copy(new_grid, grid)

	for i, row := range grid {
		for j, char := range row {
			if char == '@' {
				if check_neighbour_count(grid, i, j) < 4 {
					accesible_count += 1
					copy_row := new_grid[i]
					new_grid[i] = copy_row[:j] + string('x') + copy_row[j+1:]
				}
			}
		}
	}

	// fmt.Println()
	// for row, k := range new_grid {
	// 	fmt.Println(row, k)
	// }

	return accesible_count, new_grid

}

func check_neighbour_count(grid []string, i, j int) int {

	out := 0

	x_min := max(j-1, 0) 
	x_max := min(j+1, len(grid)-1)
	y_min := max(i-1, 0)
	y_max := min(i+1, len(grid[0])-1)


	for x := x_min; x <= x_max; x++ {
		for y := y_min; y <= y_max; y++ {
			if grid[y][x] == '@' {
				out += 1
			}
		}
	}

	// subtract 1 to account for the central square
	return out - 1
}


func load_day_4_data(filename string) []string {
	f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	out := []string{}

	for scanner.Scan() {
		row := scanner.Text()

		out = append(out, row)
	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return out
}