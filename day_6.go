package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func day_6() {

	fmt.Println("sixth day")

	// Part 1
	test_homework_p1 := load_day_6_data_p1("day_6_test.csv", 3)
	real_homework_p1 := load_day_6_data_p1("day_6.csv", 4)

	// fmt.Println(real_homework)
		

	test_out := process_part1_day6(test_homework_p1)
	fmt.Println("Part 1 test output: ", test_out)

	out := process_part1_day6(real_homework_p1)
	fmt.Println("Part 1 real output: ", out)

	// Part 2
	test_homework_p2 := load_day_6_data_p2("day_6_test.csv", 3)
	real_homework_p2 := load_day_6_data_p2("day_6.csv", 4)

	// fmt.Println(test_homework_p2)
	
	test_out_p2 := process_part1_day6(test_homework_p2)
	fmt.Println("Part 2 test output: ", test_out_p2)

	out_p2 := process_part1_day6(real_homework_p2)
	fmt.Println("Part 2 real output: ", out_p2)
}

func process_part1_day6(in []homeWorkColumn) int {

	tot := 0

	for _, column := range in {

		if column.operation == "*" {
			base := 1
			for _, mult := range column.nums {
				base *= mult
			}
			tot += base
		} else if column.operation == "+" {
			for _, sum := range column.nums {
				tot += sum
			}
		}

	}
	return tot
}

type homeWorkColumn struct {
	nums []int
	operation string
}

func load_day_6_data_p1(filename string, rows int) []homeWorkColumn {
	f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	count := 0

	out := []homeWorkColumn{}

	for scanner.Scan() {
		row := scanner.Text()

		if count == 0 {
			nums := strings.Fields(row)
			for _, num := range nums {
				num_int, nErr := strconv.Atoi(num)
				if nErr != nil {
					log.Fatal(nErr)
				}
				out = append(out, homeWorkColumn{nums: []int{num_int}})
			}

		} else if count < rows {

			nums := strings.Fields(row)
			for i, num := range nums {
				num_int, nErr := strconv.Atoi(num)
				if nErr != nil {
					log.Fatal(nErr)
				}
				out_nums := out[i].nums
				out_nums = append(out_nums, num_int)
				out[i].nums = out_nums
			}

		} else {
			ops := strings.Fields(row)
			for i, op := range ops {
				out[i].operation = op
			}
		}

		count += 1
	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return out
}

func load_day_6_data_p2(filename string, rows int) []homeWorkColumn {
	f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	row_strings := []string{}

	for scanner.Scan() {
		row := scanner.Text()
		row_strings = append(row_strings, row)
	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return process_part2_raw(row_strings)
}

func process_part2_raw(in []string) []homeWorkColumn {

	out := []homeWorkColumn{}

	operator_ind := len(in) - 1

	new_col := homeWorkColumn{nums: []int{}}

	for i := len(in[0]) - 1; i >= 0; i-- {
		num_to_add := ""
		for j := 0; j < operator_ind; j++ {
			num_to_add += string(in[j][i])
		}

		stripped := strings.Fields(num_to_add)
		// if len(stripped) == 0 {
		// 	continue
		// }
		num_to_add_int, nErr := strconv.Atoi(stripped[0])
		if nErr != nil {
			log.Fatal(nErr)
		}

		curr_nums := new_col.nums
		curr_nums = append(curr_nums, num_to_add_int)
		new_col.nums = curr_nums

		if op := in[operator_ind][i]; op != ' ' {
			new_col.operation = string(op)
			out = append(out, new_col)
			new_col = homeWorkColumn{nums: []int{}}
			i--
		}
	}


	return out

}