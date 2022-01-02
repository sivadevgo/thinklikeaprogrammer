package main

import "fmt"

func main() {
	var n int
	// fmt.Scanf("%d", &n)

	// 	for a := 0; a < n; a++ {

	// 	}
	n = 4
	s := n*2 + n*2 - 2

	for i, j := 0, 1; i < n; i, j = i+1, j+1 {
		// fmt.Print("i ", i, "j ", j)
		printSpace(i)
		printHash(j)
		// fmt.Print(s - 2*i - 2*j)
		printSpace(s - 2*i - 2*j)
		printHash(j)
		printSpace(i)
		fmt.Println()
	}

	for i, j := n-1, n; i >= 0; i, j = i-1, j-1 {
		// fmt.Print("i ", i, "j ", j)
		printSpace(i)
		printHash(j)
		// fmt.Print(s - 2*i - 2*j)
		printSpace(s - 2*i - 2*j)
		printHash(j)
		printSpace(i)
		fmt.Println()
	}
}

func printHash(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("#")
	}
}

func printSpace(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf(" ")
	}
}
