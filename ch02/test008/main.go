package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	output := []int{}

	for n >= 1 {
		v := n % 2
		output = append(output, v)
		n /= 2
	}

	print(output)
}

func print(input []int) {
	for i := len(input) - 1; i >= 0; i-- {
		fmt.Printf("%d", input[i])
	}
}
