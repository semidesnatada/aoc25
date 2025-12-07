package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day_7() {

	test_tree := load_day_7_data("day_7_test.csv")
	real_tree := load_day_7_data("day_7.csv")

	// for i, row := range test_tree {
	// 	fmt.Println(row, i)
	// }
		
	// Part 1
	test_out := process_part1_day7(test_tree)
	fmt.Println("Part 1 test output: ", test_out)

	out := process_part1_day7(real_tree)
	fmt.Println("Part 1 real output: ", out)

	// Part 2
	test_out_p2 := process_part2_day7(test_tree)
	fmt.Println("Part 2 test output: ", test_out_p2)

	out_p2 := process_part2_day7(real_tree)
	fmt.Println("Part 2 real output: ", out_p2)


}

func process_part2_day7(tree []string) int {

	int_tree := prepare_int_tree(tree)

	int_tree = recur_bfs_paths(int_tree, 0)

	// restringed := convert_int_tree_to_string(int_tree)
	// for i, row := range restringed {
	// 	fmt.Println(row, "  ", i)
	// }

	last_row := int_tree[len(int_tree)-1]
	out := 0
	for _, val := range last_row {
		if val != -1 {
			out += val
		}
	}

	return out
}

func recur_bfs_paths(int_tree [][]int, row_num int) [][]int {

	for ind, val := range int_tree[row_num] {
		if val > 0 {
			if int_tree[row_num+1][ind] == -1 {
				int_tree[row_num+1][ind-1] += val
				int_tree[row_num+1][ind+1] += val
			} else {
				int_tree[row_num+1][ind] += val
			}
		}
	}

	if row_num + 2 < len(int_tree) {
		return recur_bfs_paths(int_tree, row_num + 1)
	} else {
		return int_tree
	}
	
}

func prepare_int_tree(tree []string) [][]int {

	int_tree := [][]int{}
	var s_ind int
	for i, row := range tree {
		
		if i == 0 {
			s_ind = getInds(row, "S")[0]
		}
		new_row := []int{}
		for _, cha := range row {
			if string(cha) == "^" {
				new_row = append(new_row, -1)
			} else {
				new_row = append(new_row, 0)
			}
		}
		int_tree = append(int_tree, new_row)
	}
	// int_tree[0][s_ind] = 1
	int_tree[1][s_ind] = 1

	return int_tree

}

func convert_int_tree_to_string(int_tree [][]int) []string {

	out := []string{}

	for _, row := range int_tree {

		var new_row string

		for _, character := range row {
			if character > 0 {
				new_row += strconv.Itoa(character)
			} else if character == 0 {
				new_row += "."
			} else if character == -1 {
				new_row += "^"
			} else {
				fmt.Println(character, " error in processing ")
			}
		}

		out = append(out, new_row)

	}

	return out
}

func process_part1_day7(tree []string) int {

	out := 0

	var split int
	for i := 0; i < len(tree) -1 ; i++ {
		tree, split = processTreeStep(tree, i)
		out += split
	}

	// for i, row := range tree {
	// 	fmt.Println(row, i)
	// }

	return out
}

func processTreeStep(tree []string, row_num int) ([]string, int) {

	splits := 0

	if row_num == 0 {
		s_ind := getInds(tree[row_num], "S")
		next := tree[row_num+1]
		next = next[:s_ind[0]] + "|" + next[s_ind[0]+1:]
		tree[row_num+1] = next
	} else {
		beam_inds := getInds(tree[row_num], "|")
		next := tree[row_num+1]

		for _, b_ind := range beam_inds {
			var splitting bool
			if next[b_ind] == '^' {
				splitting = true
			} else {
				splitting = false
			}
			if splitting {
				splits += 1
				next = next[:b_ind-1] + "|" + string(next[b_ind]) + "|" + next[b_ind+2:]
			} else {
				next = next[:b_ind] + "|" + next[b_ind+1:]
			}
				
		}
		tree[row_num+1] = next
	}

	return tree, splits

}

func getInds(row string, lookup string) []int {
	// returns ascending order least of indices of locations of lookup in a row
	out := []int{}

	for i := 0; i < len(row); i++ {
		if string(row[i]) == lookup {
			out = append(out, i)
		}
	}

	return out
}

func load_day_7_data(filename string) []string {
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