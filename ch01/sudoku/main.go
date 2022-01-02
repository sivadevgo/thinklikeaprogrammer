package main

import "fmt"

func main() {
	s := getInput()
	printSudoku(s)
	fmt.Println(validate(s))

	newSudoku := generateDS(s)
	// fmt.Println(newSudoku)
	fmt.Println(validateNS(newSudoku))
	for !validateNS(newSudoku) {
		r := removeUnwantedNumbersFromRow(newSudoku)
		fmt.Println("row ", r)
		r = removeUnwantedNumbersFromColumn(newSudoku)
		fmt.Println("column ", r)
		removeUnwantedNumbersFromSquare(newSudoku)
		fmt.Println("square ", r)
	}
	printNS(newSudoku)
}

func removeUnwantedNumbersFromRow(input [][]map[int]bool) (output bool) {
	for _, cs := range input {
		for _, cm := range cs {
			if len(cm) == 1 {
				var k int
				for k = range cm {
				}
				for _, m := range cs {
					_, ok := m[k]
					if ok && len(m) != 1 {
						delete(m, k)
						output = true
					}
				}
			}
		}
	}
	return
}

func removeUnwantedNumbersFromColumn(input [][]map[int]bool) (output bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cm := input[j][i]
			if len(cm) == 1 {
				var k int
				for k = range cm {
				}
				for a := 0; a < 9; a++ {
					m := input[a][i]
					_, ok := m[k]
					if ok && len(m) != 1 {
						delete(m, k)
						output = true
					}
				}
			}

		}
	}
	return
}

func removeUnwantedNumbersFromSquare(input [][]map[int]bool) (output bool) {
	for i := 0; i < len(input); i += 3 {
		for j := 0; j < len(input); j += 3 {
			for a := i; a < i+3; a++ {
				for b := j; b < j+3; b++ {
					cm := input[a][b]
					if len(cm) == 1 {
						var k int
						for k = range cm {
						}
						for a2 := i; a2 < i+3; a2++ {
							for b2 := j; b2 < j+3; b2++ {
								m := input[a2][b2]
								_, ok := m[k]
								if ok && len(m) != 1 {
									delete(m, k)
									output = true
								}
							}
						}
					}
				}
			}
		}
	}
	return
}

func generateDS(input [][]int) [][]map[int]bool {
	output := make([][]map[int]bool, 9)
	for i, cs := range input {
		currentSlice := make([]map[int]bool, 9)
		output[i] = currentSlice
		for j, ce := range cs {
			if ce == 0 {
				output[i][j] = map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}
			} else {
				output[i][j] = map[int]bool{ce: true}
			}
		}
	}
	return output
}

func printSudoku(input [][]int) {
	for _, currentLine := range input {
		fmt.Println(currentLine)
	}
}

func getInput() [][]int {
	return [][]int{
		{0, 9, 1, 0, 6, 0, 7, 0, 0},
		{0, 0, 0, 0, 8, 2, 0, 3, 9},
		{5, 0, 3, 0, 0, 0, 2, 0, 0},
		{0, 0, 0, 9, 1, 3, 0, 6, 2},
		{0, 0, 2, 4, 0, 6, 8, 0, 0},
		{1, 4, 0, 8, 2, 5, 0, 0, 0},
		{0, 0, 9, 0, 0, 0, 5, 0, 7},
		{6, 7, 0, 1, 5, 0, 0, 0, 0},
		{0, 0, 5, 0, 4, 0, 6, 9, 0},
	}
}

func printNS(input [][]map[int]bool) {

	for _, cs := range input {
		for _, ce := range cs {
			if len(ce) != 1 {
				fmt.Println("********** ERROR ************")
				return
			}
		}
	}

	s := make([][]int, 9)
	for i, cs := range input {
		currentSlice := make([]int, 9)
		s[i] = currentSlice
		for j, ce := range cs {
			for k := range ce {
				s[i][j] = k
			}
		}
	}
	printSudoku(s)
}

func validateNS(input [][]map[int]bool) bool {

	for _, cs := range input {
		for _, ce := range cs {
			if len(ce) != 1 {
				return false
			}
		}
	}

	s := make([][]int, 9)
	for i, cs := range input {
		currentSlice := make([]int, 9)
		s[i] = currentSlice
		for j, ce := range cs {
			for k := range ce {
				s[i][j] = k
			}
		}
	}
	return validate(s)
}

func validate(input [][]int) bool {

	// fmt.Println("Starting row wise")
	//validate rowwise
	for _, cs := range input {
		map1 := make(map[int]bool)
		for _, cn := range cs {
			map1[cn] = true
		}
		if len(map1) != 9 {
			return false
		}
	}
	fmt.Println("Starting column wise")

	//validate columnwise
	for i := 0; i < len(input); i++ {
		map1 := make(map[int]bool)
		for j := 0; j < len(input); j++ {
			map1[input[j][i]] = true
		}
		if len(map1) != 9 {
			// fmt.Println(map1)
			return false
		}
	}

	fmt.Println("Starting square wise")
	for i := 0; i < len(input); i += 3 {
		for j := 0; j < len(input); j += 3 {
			map1 := make(map[int]bool)
			for a := i; a < i+3; a++ {
				for b := j; b < j+3; b++ {
					map1[input[a][b]] = true
					// fmt.Print(" ", a, ",", b)
				}
				// fmt.Println()
			}
			if len(map1) != 9 {
				fmt.Println(false)
				return false
			}
		}
	}
	return true
}
