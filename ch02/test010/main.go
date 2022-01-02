package main

import "fmt"

func main() {
	input := []int{}
	var source int
	var target int
	var c byte
	fmt.Println("Enter source")
	fmt.Scanf("%d\n", &source)
	fmt.Println("Enter target")
	fmt.Scanf("%d\n", &target)

	fmt.Println("Enter input")
	fmt.Scanf("%c", &c)
	for c >= '0' && c <= '9' {
		input = append(input, int(c)-'0')
		fmt.Scanf("%c", &c)
	}
	// fmt.Println("input ", input)

	output := convertSourceToTarget(input, source, target)
	print(output)
}

func print(input []int) {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] >= 10 && input[i] <= 15 {
			fmt.Printf("%c", 'A'+input[i]-10)
		} else {
			fmt.Print(input[i])
		}
	}
	fmt.Println()
}

func convertSourceToTarget(input []int, source, target int) []int {
	output := []int{}
	var sum int
	for _, v := range input {
		sum *= source
		sum += v
		// fmt.Println("s", sum)
	}
	// fmt.Println("sum=", sum)
	for sum != 0 {
		output = append(output, sum%target)
		sum = sum / target
	}

	return output
}
