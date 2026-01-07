package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func day_11(print bool) {


	// fmt.Println("eleventh day")

	test_racks_p1 := load_day_11_data("day_11_test.csv")
	test_racks_p2 := load_day_11_data("day_11_test_p2.csv")
	real_racks := load_day_11_data("day_11.csv")

	// Part 1
	test_out := process_part1_day11(test_racks_p1)
	out := process_part1_day11(real_racks)

	if print {
		fmt.Println()
		fmt.Println()
		fmt.Println("** Day 11 **")
		fmt.Println("// Part 1 //")
		fmt.Println("Part 1 test output: ", test_out)
		fmt.Println("Part 1 real output: ", out)
	}

	// Part 2
	test_out_p2 := process_part2_day11(test_racks_p2)
	// can't achieve part 2 witht just a simple dfs.
	// need ot either prune the search space, 
	// or think of an alternative approach
	out_p2 := process_part2_day11(real_racks)

	if print {
		fmt.Println()
		fmt.Println("// Part 2 //")
		fmt.Println("Part 2 test output: ", test_out_p2)
		fmt.Println("Part 2 real output: ", out_p2)
	}
	
}

func process_part2_day11(in map[string][]string) int {

	// no routes from dac to fft
	// so all solutions are svr -> ... -> fft -> ... -> dac -> ... -> out

	// get all paths from start to fft. as there are no paths from dac to fft, 
	// then all valid paths must use one of paths from start to fft
	// svr_fft_paths, ignores := get_paths(in, "svr", "fft")

	// previous iteration kept track of "ignores", an arrary which
	//  contains all of the nodes which lead to loops, so we can safely
	//  ignore nodes where all of their children result in either missed opportunities or loops
	//  for the puzzle input, however, the search space quickly blew up such that 
	//  we couldn't ignore any + there were no loops! So removed for this version

	g := graph{
		cache: make(map[string]int),
		path_map: in,
	}

	svr_fft := g.cached_dfs_part_2("svr", "fft")
	fft_dac := g.cached_dfs_part_2("fft", "dac")
	dac_out := g.cached_dfs_part_2("dac", "out")

	// still calculating these anyway for completeness, should = 0
	svr_dac := g.cached_dfs_part_2("svr", "dac")
	dac_fft := g.cached_dfs_part_2("dac", "fft")
	fft_out := g.cached_dfs_part_2("fft", "out")

	total_path_1 := svr_fft * fft_dac * dac_out
	total_path_2 := svr_dac * dac_fft * fft_out

	return total_path_1 + total_path_2
}

func process_part1_day11(in map[string][]string) int {

	// fmt.Println(in)

	route := make(map[string]int)
	route["you"] = 0

	paths_tracker := [][]string{}

	paths_tracker = dfs_for_out(route, "you", "out", in, paths_tracker)

	// fmt.Println(paths_tracker)

	return len(paths_tracker)
}

type graph struct {
	cache map[string]int
	path_map map[string][]string
}

func (g *graph) cached_dfs_part_2(start, target string) int {
	// simplified dfs to cache output and only consider
	// the number of paths rather than generate each path
	// this is safe becausue we have previously identified
	// that there are no loops in the graph (i.e. len(ignore)==0)
	// so path history is irrelevant.

	if cached_val, ok := g.cache[start+target]; ok {
		return cached_val
	}

	out := 0

	for _, next := range g.path_map[start] {
			if next == target {
				out += 1
				continue
			}

			out_add := g.cached_dfs_part_2(next, target)
			g.cache[next+target] = out_add
			out += out_add

	}

	g.cache[start+target] = out

	return out
}

func dfs_for_out_p2(
		current_route map[string]int,
		current, target string,
		path_map map[string][]string, 
		out_tracker [][]string, 
		ignore *map[string]bool,
	) [][]string {

		// Thiis function is now redundant as it was not fast enough.
		// we need to cache function outputs.
		// we can cache easily as we do not actually care about storing 
		// path history because there are no loops in this graph
		// (as proved separately)

	next_ind := len(current_route)

	var found_here, found_below bool

	//  *** Next test *** //
	// number of ignorable coords seems to top off at around 221.
	// maybe removing them completely from the input would speed up runtime.

	// fmt.Println(len(*ignore))

	// _, ok_d := (*ignore)["dac"]
	// _, ok_f := (*ignore)["fft"]

	// fmt.Println(ok_d, ok_f)

	for _, next := range path_map[current] {

		// if depth <= 3 {
		// 	fmt.Println("doiung a new one at depth ", depth)
		// }

		if _, ok := (*ignore)[next]; ok {
			continue
		}

		if _, ok := current_route[next]; !ok {
			new_current_map := make(map[string]int)
			for key, val := range current_route {
				new_current_map[key] = val
			}
			new_current_map[next] = next_ind
			
			if next == target {
				found_here = true
				new_route := make([]string, len(new_current_map))

				for key, val := range new_current_map {
					new_route[val] = key
				}

				out_tracker = append(out_tracker, new_route)
				continue
			}
			curr_len := len(out_tracker)
			out_tracker = dfs_for_out_p2(new_current_map, next, target, path_map, out_tracker, ignore)
			
			if len(out_tracker) > curr_len {
				found_below = true
			}
		}
		
	}

	if !found_here && !found_below {
		(*ignore)[current] = true
	}

	return out_tracker

}

func dfs_for_out(current_route map[string]int, current, target string, path_map map[string][]string, out_tracker [][]string) [][]string {

	next_ind := len(current_route)

	for _, next := range path_map[current] {

		// for part 2, maybe we can check if any routes are found from a node.
		// if not, then we can permanently ignore the node, should it appear again.
		// note: the above was the origin of the "ignore" argument of now redundant function rewrites.

		if _, ok := current_route[next]; !ok {
			new_current_map := make(map[string]int)
			for key, val := range current_route {
				new_current_map[key] = val
			}
			new_current_map[next] = next_ind
			
			if next == target {
				new_route := make([]string, len(new_current_map))

				for key, val := range new_current_map {
					new_route[val] = key
				}

				out_tracker = append(out_tracker, new_route)
				continue
			}

			out_tracker = dfs_for_out(new_current_map, next, target, path_map, out_tracker)
			
		}
		// if len(out_tracker) >= 1 {
		// 	fmt.Println(out_tracker)
		// }
		
	}

	return out_tracker

}


func load_day_11_data(filename string) map[string][]string {
		f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	out := make(map[string][]string)

	for scanner.Scan() {
		
		row := scanner.Text()

		parts := strings.Fields(row)
		prime := strings.Split(parts[0], ":")

		out[prime[0]] = make([]string,len(parts[1:]))

		for i, val := range parts[1:] {
			out[prime[0]][i] = val
		}
		
	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return out
}


// func identify_loop_starts(in map[string][]string) map[string][]string {

// 	// this has identified there are no loops in the graph

// 	new_map := make(map[string][]string)

// 	route := make(map[string]int)
// 	route["svr"] = 0

// 	paths_tracker := [][]string{}

// 	ignore := make(map[string]bool)

// 	paths_tracker = dfs_for_out_p2(route, "svr", "dac", in, paths_tracker, &ignore)


// 	return new_map
// }

// func get_paths(graph map[string][]string, start, end string) ([][]string, map[string]bool) {

// 	// redundant helper function made for part 2 which is no longer useful

// 	route := make(map[string]int)
// 	route[start] = 0

// 	paths_tracker := [][]string{}

// 	ignore := make(map[string]bool)

// 	paths_tracker = dfs_for_out_p2(route, start, end, graph, paths_tracker, &ignore, 1)

// 	// fmt.Println(len(ignore))

// 	// for i, path := range paths_tracker{
// 	// 	fmt.Println(i, path)
// 	// }

// 	return paths_tracker, ignore
// }
