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

func day_9() {

	// super slow - haven't identified bottleneck

	fmt.Println("ninth day")

	test_tiles := load_day_9_data("day_9_test.csv")
	real_tiles := load_day_9_data("day_9.csv")

	// for i, tile := range test_tiles {
	// 	fmt.Println(tile, i)
	// }

	// Part 1
	test_out := process_part1_day9(test_tiles)
	fmt.Println("Part 1 test output: ", test_out)

	out := process_part1_day9(real_tiles)
	fmt.Println("Part 1 real output: ", out)

	// Part 2
	test_out_p2 := process_part2_day9(test_tiles)
	fmt.Println("Part 2 test output: ", test_out_p2)

	out_p2 := process_part2_day9(real_tiles)
	fmt.Println("Part 2 real output: ", out_p2)

}

func process_part2_day9(coords []coord) int {

	// edge_map := get_edge_map(coords)

	rects := []rectangle{}

	for i := 0; i < len(coords) - 1; i++ {
		for j := i + 1; j < len(coords); j++ {
			rects = append(rects, rectangle{
				c1: coords[i],
				c2: coords[j],
			})
		}
	}

	slices.SortFunc(rects, func(a, b rectangle) int {
		aa := a.area()
		ab := b.area()
		if aa > ab {return -1}
		if ab > aa {return 1}
		return 0
	})

	shappy := createShape(coords)

	max_area := 0

	for _, rect := range rects {
		// fmt.Println(rect, rect.area())
		// if i % 1000 == 0 {
		// 	fmt.Println("Checked ", i, " rects of ", len(rects))
		// }

		// note that the below inShape function calls are not necessary,
		// and are only included to try to speed up computation
		
		// if rect corners are not in shape, then the rect is not in shape
		// a, b := rect.getOtherCoords()
		// // fmt.Println(a, b)
		// if !a.inShape(shappy) {
		// 	// fmt.Println("failed at corner a point", a)
		// 	continue
		// }
		// if !b.inShape(shappy) {
		// 	// fmt.Println("failed at corner b point", b)
		// 	continue
		// }

		// // check if a point in the rect is in the shape
		// test_x := min(rect.c1.x, rect.c2.x)+1
		// test_y := min(rect.c1.y, rect.c2.y)+1
		// test_coord := coord{x: test_x, y:test_y}
		// if !test_coord.inShape(shappy) {
		// 	// the point inside the rect is not inside the shape, so the rect is not inside the shape
		// 	// fmt.Println("failed at in shape point")
		// 	continue
		// }

		// if any edge of the shape intersects with the rectangle, then the rectangle is not valid
		// so check whether any point on a horizontal edge of the rect is in the vertical edges of the shape
		// and check whether any point on a vertical edge of the rect is in the horizontal edges of the shape
		if rect.doEdgesIntersectWithShape(shappy) {
			// fmt.Println("failed at edge intersection point")
			continue
		}

		max_area = rect.area()
		// fmt.Println(i)
		break
	}

	return max_area

}

func createShape(coords []coord) shape {

	coord_map := make(map[coord]corner_dir)

	vert_edges := make(map[int]map[int]bool)
	hori_edges := make(map[int]map[int]bool)

	last_coord := coords[len(coords)-1]
	next_coord := coords[1]
	
	for i, coo := range coords {

		hori_d := coo.x - last_coord.x
		vert_d := coo.y - last_coord.y

		if hori_d == 0 {
			if vert_d > 0 {
				for j := last_coord.y + 1; j < coo.y; j++ {
					if _, ok := vert_edges[j]; ok {
						vert_edges[j][coo.x] = true
					} else {
						vert_edges[j] = make(map[int]bool)
						vert_edges[j][coo.x] = true
					}
					coord_map[coo] = corner_dir{
						u: false,
						d: true,
						l: next_coord.x < coo.x,
						r: next_coord.x > coo.x,
					}	
				}

			} else {
				for j := last_coord.y - 1; j > coo.y; j-- {
					if _, ok := vert_edges[j]; ok {
						vert_edges[j][coo.x] = true
					} else {
						vert_edges[j] = make(map[int]bool)
						vert_edges[j][coo.x] = true
					}
					coord_map[coo] = corner_dir{
						u: true,
						d: false,
						l: next_coord.x < coo.x,
						r: next_coord.x > coo.x,
					}
				}
			}
		} else {
			if hori_d > 0 {
				for j := last_coord.x + 1; j < coo.x; j++ {
					if _, ok := hori_edges[j]; ok {
						hori_edges[j][coo.y] = true
					} else {
						hori_edges[j] = make(map[int]bool)
						hori_edges[j][coo.y] = true
					}
					coord_map[coo] = corner_dir{
						u: next_coord.y < coo.y,
						d: next_coord.y > coo.y,
						l: false,
						r: true,
					}
				}
			} else {

				for j := last_coord.x - 1; j > coo.x; j-- {
					if _, ok := hori_edges[j]; ok {
						hori_edges[j][coo.y] = true
					} else {
						hori_edges[j] = make(map[int]bool)
						hori_edges[j][coo.y] = true
					}
					coord_map[coo] = corner_dir{
						u: next_coord.y < coo.y,
						d: next_coord.y > coo.y,
						l: true,
						r: false,
					}
				}
			}
		}
		last_coord = coo
		if i + 2 >= len(coords) {
			next_coord = coords[(i+2) % len(coords)]
		} else {
			next_coord = coords[i+2]
		}
	}

	return shape{
		horiz_edges: hori_edges,
		verti_edges: vert_edges,
		corners: coord_map,
	}
}

func (r rectangle) doEdgesIntersectWithShape(s shape) bool {
	// checks whether a ring internal to the rectangle 
	// intersects with any perpendicular lines of the shape

	min_x := min(r.c1.x, r.c2.x)
	max_x := max(r.c1.x, r.c2.x)
	
	min_y := min(r.c1.y, r.c2.y)
	max_y := max(r.c1.y, r.c2.y)

	for xx := min_x + 1; xx < max_x; xx++ {
		if _, ok := s.corners[coord{x: xx, y: min_y+1}]; ok {
			return true
		}
		if _, ok := s.corners[coord{x: xx, y: max_y-1}]; ok {
			return true
		}
		if _, ok := s.verti_edges[min_y+1][xx]; ok {
			return true
		}
		if _, ok := s.verti_edges[max_y-1][xx]; ok {
			return true
		}
	}

	for yy := min_y + 1; yy < max_y; yy++ {
		if _, ok := s.corners[coord{x: min_x+1, y: yy}]; ok {
			return true
		}
		if _, ok := s.corners[coord{x: max_x-1, y: yy}]; ok {
			return true
		}
		if _, ok := s.horiz_edges[min_x+1][yy]; ok {
			return true
		}
		if _, ok := s.horiz_edges[max_x-1][yy]; ok {
			return true
		}
	}
	return false
}

func (r rectangle) getOtherCoords() (coord, coord) {
	return coord{x: r.c1.x, y: r.c2.y}, coord{x: r.c2.x, y: r.c1.y}
}

func (r rectangle) area() int {
	x_len := r.c1.x - r.c2.x
	x_len = max(x_len, -x_len) + 1
	y_len := r.c1.y - r.c2.y
	y_len = max(y_len, -y_len) + 1

	return x_len * y_len
}

func (c coord) inShape(s shape) bool {

	// check if it's a corner
	if _, ok := s.corners[c]; ok {
		return true
	}

	// check if it's an edge
	if s.horiz_edges[c.x][c.y] || s.verti_edges[c.y][c.x] {
		return true
	}

	// check if it's within the shape
	verti_lines, horiz_lines := 0, 0

	var ud_sum, lr_sum int
	for i := c.x; i >= 0; i-- {
		if _, ok := s.verti_edges[c.y][i]; ok {
			verti_lines += 1
		} else if cc, ok := s.corners[coord{x: i, y: c.y}]; ok {
			if cc.u {
				ud_sum ++
			} else if cc.d {
				ud_sum --
			}
		}
	}
	verti_lines += ud_sum / 2

	for i := c.y; i >= 0; i-- {
		if _, ok := s.horiz_edges[c.x][i]; ok {
			horiz_lines += 1
		} else if cc, ok := s.corners[coord{x: c.x, y: i}]; ok {
			if cc.l {
				lr_sum ++
			} else if cc.r {
				lr_sum --
			}
		}
	}
	horiz_lines += lr_sum / 2

	// fmt.Println("vertis: ", verti_lines)
	// fmt.Println("horizs: ", horiz_lines)

	if verti_lines % 2 == 0 {
		return false
	}
	if horiz_lines % 2 == 0 {
		return false
	}
	return true

}

type shape struct {
	// map where the first key is the x coord, yielding a map which has keys that are y coords which are on edges
	horiz_edges map[int]map[int]bool
	// map where the first key is the y coord, yielding a map which has key that are x coords which are on edges
	verti_edges map[int]map[int]bool
	corners map[coord]corner_dir
}

type corner_dir struct {
	u, d, l, r bool
}

type rectangle struct {
	c1, c2 coord		
}

type coord struct {
	x, y int
}

func load_day_9_data(filename string) []coord {
	f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	out := []coord{}


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

		out = append(out, coord{
			x: x_i,
			y: y_i,
		})

	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return out
}

func process_part1_day9(coords []coord) int {

	max_size := 0

	for i := 0; i < len(coords) - 1; i++ {
		for j := i + 1; j < len(coords); j++ {
			x_len := (coords[i].x - coords[j].x + 1)
			x_len = max(x_len, -x_len)
			y_len := (coords[i].y - coords[j].y + 1)
			y_len = max(y_len, -y_len)
			area := x_len * y_len
			if area > max_size {
				max_size = area
			}
		}
	}

	return max_size

}