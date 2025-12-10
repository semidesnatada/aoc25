package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
)

func day_10() {


	// this solution requires an external ILP solver as
	// I could not be bothered to write my own.

	fmt.Println("tenth day")

	test_machines := load_day_10_data("day_10_test.csv")
	real_machines := load_day_10_data("day_10.csv")

	// for i, machine := range test_machines {
	// 	fmt.Println(i)
	// 	fmt.Println(machine.successState)
	// 	fmt.Println(machine.buttons)
	// 	fmt.Println(machine.joltage)
	// }

	// Part 1
	test_out := process_part1_day10(test_machines)
	fmt.Println("Part 1 test output: ", test_out)

	out := process_part1_day10(real_machines)
	fmt.Println("Part 1 real output: ", out)

	// Part 2
	test_out_p2 := process_part2_day10(test_machines)
	fmt.Println("Part 2 test output: ", test_out_p2)

	out_p2 := process_part2_day10(real_machines)
	fmt.Println("Part 2 real output: ", out_p2)

}

func process_part2_day10(machines []machine) int {
	out := 0

	for _, mac := range machines {
		// fmt.Println(i, " of ", len(machines))

		// linear programming solution.
		// uses golp bindings for LPSolve solver
		A, b := formulate_lp(mac)
		if val := solve_lp(A, b); val != -1 {
			out += val
		} else {
			log.Fatal("couldn't solve one of the problems")
		}

		// this is for the multicomb approach, which does not work with the real input size
		// fmt.Println(mac.successState, "can be solved with ", test_machine_p2(mac) , " presses.")
		// if x := test_machine_p2(mac); x != -1 {
		// 	out += x
		// } else {
		// 	log.Fatal("big error")
		// }

	}

	return out

}

func process_part1_day10(machines []machine) int {
	out := 0

	for _, mac := range machines {

		// fmt.Println(mac.successState, "can be solved with ", test_machine(mac) , " presses.")
		if x := test_machine(mac); x != -1 {
			out += x
		} else {
			log.Fatal("big error")
		}

	}

	return out

}

func test_machine(mac machine) int {
	n := len(mac.buttons)

	for k := 0; k <= n; k++ {
		button_combinations := generateKCombinations(n, k)
		for _, button_combination := range button_combinations {
			if test_button_combination(mac, button_combination) {
				return k
			}
		}
	}
	return -1
}

func test_machine_p2(mac machine) int {
	// for the multicomb approach, which does not work with the puzzle input
	n := len(mac.buttons)

	var min_presses int
	var max_presses int
	for _, succ := range mac.joltage {
		min_presses = max(min_presses, succ)
		max_presses += succ
	}

	for k := min_presses; k <= max_presses; k++ {
		fmt.Println("generating multicombs for: n = ", n, " and k = ", k)
		button_combinations := generateKMulticombinations(n, k)
		fmt.Println(button_combinations)
		for i, button_combination := range button_combinations {
			fmt.Println(i, " of ", len(button_combinations))
			if test_button_combination_p2(mac, button_combination) {
				return k
			}
		}
	}
	return -1
}

func test_button_combination_p2(mac machine, button_combination []int) bool {
	start := make([]int, len(mac.successState))

	for _, ind := range button_combination {
		to_press := mac.buttons[ind]
		for but := range to_press {
			start[but] ++
		}	
	}

	for i, val := range start {
		if val != mac.joltage[i] {
			return false
		}
	}

	return true
}

func test_button_combination(mac machine, button_combination []int) bool {
	start := make([]bool, len(mac.successState))

	for _, ind := range button_combination {
		to_press := mac.buttons[ind]
		for but := range to_press {
			start[but] = !start[but]
		}	
	}

	for i, val := range start {
		if val != mac.successState[i] {
			return false
		}
	}

	return true
}

type button map[int]bool

type machine struct {
	successState []bool
	buttons []button
	joltage []int
}

func load_day_10_data(filename string) []machine {
		f, err := os.Open(fmt.Sprintf("inputs/%s", filename))
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	
	scanner := bufio.NewScanner(f)

	out := []machine{}

	for scanner.Scan() {
		
		row := scanner.Text()
		
		new_machine := machine{buttons: []button{}}

		parts := strings.Fields(row)
		for i, part := range parts {
			if i == 0 {
				b := make([]bool, len(part)-2)
				for i := 0; i < len(part) - 2; i++ {
					if part[i+1] == '.' {
						b[i] = false
					} else if part[i+1] == '#' {
						b[i] = true
					} else {
						log.Fatal("something went wrong")
					}
				}
				new_machine.successState = b
			} else if i == len(parts) - 1 {
				jolts_s := strings.Split(part[1:len(part)-1],",")
				jolts_i := make([]int, len(jolts_s))
				for j, jolt := range jolts_s {
					ji, jErr := strconv.Atoi(jolt)
					if jErr != nil {
						log.Fatal(jErr)
					}
					jolts_i[j] = ji
				}
				new_machine.joltage = jolts_i
			} else {
				butts_s := strings.Split(part[1:len(part)-1],",")
				butts_i := make(map[int]bool)
				for _, butt := range butts_s {
					ji, jErr := strconv.Atoi(butt)
					if jErr != nil {
						log.Fatal(jErr)
					}
					butts_i[ji] = true
				}
				new_machine.buttons = append(new_machine.buttons, butts_i)
			}
		}

		out = append(out, new_machine)

	}
	
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return out
}

func formulate_lp(mac machine) ([][]float64, []float64) {

	A := make([][]float64, len(mac.joltage))
	for i := range A {
		A[i] = make([]float64, len(mac.buttons))
	}

	for i, button := range mac.buttons {
		for row := range button {
			A[row][i] = 1
		}
	}

	b := make([]float64, len(mac.joltage))

	for i, target := range mac.joltage {
		b[i] = float64(target)
	}

	return A, b
}

func solve_lp(A [][]float64, b []float64) int {

	// solves the integer linear programming 
	// problem where we want to minimise the sum of the
	// output vector elements.

	// uses the golp bindings github.com/draffensperger/golp
	// and requires the LPSolve library to be installed https://lpsolve.sourceforge.net/5.5/

	n := len(A[0])
    lp := golp.NewLP(0, n)

    // set the objective functiton to minimize sum of x - this is minimised by default
    obj := make([]float64, n)
    for i := range obj {
        obj[i] = 1
    }
	lp.SetObjFn(obj)

    // add the constraints that (sum over j) A[i][j] * x[j] = b[i]
    for i := range A {
        lp.AddConstraint(A[i], golp.EQ, b[i])
    }

    // force the solver to use integers only. by default the solver uses positive values
    for j := 0; j < n; j++ {
        lp.SetInt(j, true)
	}

    result := lp.Solve()
    if result != 0 {
        return -1
    }

    solution := lp.Variables()
    // fmt.Println("x =", solution)

	sum := 0
	for _, val := range solution {
		sum += int(val)
	}

	return sum

}

func generateKMulticombinations(n, k int) [][]int {

	// not appropriate for the puzzle input, leads to calculations like 64 Choose 58
	// which are far too large to compute in reasonable time

	// relies on the fact that there is a bijective mapping 
	// between the K-Com (n+k-1)C(k) and the Multi-K-comb of n choose k
	// so to generate combination from a set with repetition, we just need to find
	// a K combination of a larger set and map the output to our correct size
	// https://en.wikipedia.org/wiki/Combination#Number_of_combinations_with_repetition

	out := generateKCombinations(n+k-1, k)

	for i, comb := range out {
		mapped_multi := make([]int, len(comb))
		for j, val := range comb {
			mapped_multi[j] = val - j
		}
		out[i] = mapped_multi
	}

	return out
}

func generateKCombinations(n, k int) [][]int {

	// function to generate k combinations adapted from
	//  https://www.geeksforgeeks.org/dsa/make-combinations-size-k/

	out := [][]int{}
	temp := []int{}

	combine(&out, &temp, n, 1, k)

	return out
}

func combine(out *[][]int, temp *[]int, n, start, k int) *[][]int {

	if k == 0 {
		// fmt.Println(*temp)
		new_temp := make([]int, len(*temp))
		copy(new_temp, *temp)
		// fmt.Println(new_temp, *temp)
		*out = append(*out, new_temp)
		
		return out
	}

	for i := start; i <= n; i++ {
		*temp = append(*temp, i-1)
		combine(out, temp, n, i+1, k-1)
		*temp = (*temp)[:len(*temp)-1]
	}

	return nil

}