package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day_8() {

	test_juncs := load_day_8_data("day_8_test.csv")
	real_tree := load_day_8_data("day_8.csv")

	// for _, row := range test_juncs {
	// 	fmt.Println(row)
	// }
		
	// Part 1
	test_out := process_part1_day8(test_juncs, 10)
	fmt.Println("Part 1 test output: ", test_out)

	out := process_part1_day8(real_tree, 1000)
	fmt.Println("Part 1 real output: ", out)

	// Part 2
	test_out_p2 := process_part2_day8(test_juncs)
	fmt.Println("Part 2 test output: ", test_out_p2)

	out_p2 := process_part2_day8(real_tree)
	fmt.Println("Part 2 real output: ", out_p2)


}

func process_part2_day8(js []junctionBox) int {

	links := getSortedDists(js)

	loops := []map[int]map[int]bool{}

	links_added := 0

	for i := 0; i < len(links); i++ {
		// iterate through the links. can either: 
		// merge two existing loops, if each j is already in different loops
		// create a new loop, if neither j is in any loop
		// add one j to an existing loop
		// do nothing if these js are already in the same loop
		new_link := links[i]

		// fmt.Println(new_link)

		// truth := true
		loops, _ = process_new_link(loops, new_link.j1.id, new_link.j2.id)

		slices.SortFunc(loops, func(a, b map[int]map[int]bool) int {
			if len(a) > len(b) {
				return -1
			}
			if len(b) > len(a) {
				return 1
			}
			return 0
		})

		links_added ++

		if len(loops[0]) == len(js) {
			return new_link.j1.x * new_link.j2.x
		}
		
	}
	
	// for _, loop := range loops {
	// 	for ind, vals := range loop {
	// 		fmt.Println(ind, vals)
	// 	}
	// }

	// return len(loops[0])
	return 0
}

func process_part1_day8(js []junctionBox, n int) int {

	links := getSortedDists(js)

	loops := []map[int]map[int]bool{}

	links_added := 0

	for i := 0; i < len(links); i++ {
		// iterate through the links. can either: 
		// merge two existing loops, if each j is already in different loops
		// create a new loop, if neither j is in any loop
		// add one j to an existing loop
		// do nothing if these js are already in the same loop
		new_link := links[i]

		// fmt.Println(new_link)

		// truth := true
		loops, _ = process_new_link(loops, new_link.j1.id, new_link.j2.id)

		slices.SortFunc(loops, func(a, b map[int]map[int]bool) int {
			if len(a) > len(b) {
				return -1
			}
			if len(b) > len(a) {
				return 1
			}
			return 0
		})

		// fmt.Println(truth)
		// if truth {
		links_added ++
		// }
		// fmt.Printf("completed another loop - added %d links\n", links_added)
		// fmt.Println()
		if links_added >= n {
			break
		}
		
	}
	
	// for i, loop := range loops {
	// 	fmt.Println(i, loop)
	// }

	return len(loops[0]) * len(loops[1]) * len(loops[2])
}

func process_new_link(loops []map[int]map[int]bool, id_1, id_2 int) ([]map[int]map[int]bool, bool) {

	// returns true and mutated input, if it is mutated
	// returns false and unmutated input otherwise

	for loop_ind_1, loop_1 := range loops {

		if _, ok_1 := loop_1[id_1]; ok_1 {

			// found id_1 in loop_1

			if _, ok_2 := loop_1[id_2]; ok_2 {
				// ids are already in the same loop
				// so do nothing
				return loops, false
			}

			for loop_ind_2, loop_2 := range loops {

				if _, ok_2 := loop_2[id_2]; ok_2 {
					// found map containing the other id
					// merge the loops with loop_ind_1 and loop_ind_2

					// merge the two existing maps
					new_m := merge_maps(loops[loop_ind_1], loops[loop_ind_2])
					// add new ids to merged map
					new_m[id_1][id_2] = true
					new_m[id_2][id_1] = true

					if loop_ind_2 < loop_ind_1 {
						loops = slices.Delete(loops, loop_ind_1,  loop_ind_1+1)
						loops = slices.Delete(loops, loop_ind_2,  loop_ind_2+1)
					} else {
						loops = slices.Delete(loops, loop_ind_2,  loop_ind_2+1)
						loops = slices.Delete(loops, loop_ind_1,  loop_ind_1+1)
					}

					loops = append(loops, new_m)

					return loops, true
				}

			}

			// couldn't find a loop containing the other id
			// so add it to the loop we have found it in
			
			loops[loop_ind_1][id_1][id_2] = true
			new_m2 := make(map[int]bool)
			new_m2[id_1] = true
			loops[loop_ind_1][id_2] = new_m2

			return loops, true

		}

	}

	// couldn't find first id in any loops
	// check if second id exists in others, if not  
	// add new loop to the list

	for loop_ind_1, loop := range loops {
		if _, ok_2 := loop[id_2]; ok_2 {
			// second id does exist, so just add id to that map

			loops[loop_ind_1][id_2][id_1] = true
			new_m := make(map[int]bool)
			new_m[id_2] = true
			loops[loop_ind_1][id_1] = new_m
			return loops, true
		}
	}

	new_loop := make(map[int]map[int]bool)

	new_1 := make(map[int]bool)
	new_1[id_2] = true
	new_2 := make(map[int]bool)
	new_2[id_1] = true

	new_loop[id_1] = new_1
	new_loop[id_2] = new_2

	loops = append(loops, new_loop)

	return loops, true
}

func merge_maps(m1, m2 map[int]map[int]bool) map[int]map[int]bool {

	m_out := make(map[int]map[int]bool)

	for key, val := range m1 {
		m_out[key] = val
	}
	for key, val := range m2 {
		m_out[key] = val
	}

	return m_out
}

func getSortedDists(js []junctionBox) []linkedJunction {

	links := []linkedJunction{}

	for i := 0; i < len(js) - 1; i++ {
		for j := i+1; j < len(js); j++ {

			links = append(links, linkedJunction{
				j1: js[i],
				j2: js[j],
				dist: dist_3d(js[i], js[j]),
			})

		}
	}

	slices.SortFunc(links, 
		func(a, b linkedJunction) int {
			if a.dist > b.dist {
				return 1
			}
			if a.dist < b.dist {
				return -1
			}
			return 0
		})

	return links

}

func dist_3d(j1, j2 junctionBox) float64 {

	x_0 := math.Pow(float64(j1.x-j2.x), 2)
	y_0 := math.Pow(float64(j1.y-j2.y), 2)
	z_0 := math.Pow(float64(j1.z-j2.z), 2)

	d_3 := x_0 + y_0 + z_0

	return math.Pow(d_3, 1.0/3.0)

}

type linkedJunction struct {
	j1, j2 junctionBox
	dist float64
}

type junctionBox struct {
	id, x, y, z int
}

func load_day_8_data(filename string) []junctionBox {
	f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	out := []junctionBox{}

	count := 0

	for scanner.Scan() {
		
		row := scanner.Text()

		coord_s := strings.Split(row, ",")
		x_i, xErr := strconv.Atoi(coord_s[0])
		if xErr != nil {
			log.Fatal(xErr)
		}
		y_i, yErr := strconv.Atoi(coord_s[1])
		if yErr != nil {
			log.Fatal(yErr)
		}
		z_i, zErr := strconv.Atoi(coord_s[2])
		if zErr != nil {
			log.Fatal(zErr)
		}

		out = append(out, junctionBox{
			id: count,
			x: x_i,
			y: y_i,
			z: z_i,
		})
		count ++

	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return out
}