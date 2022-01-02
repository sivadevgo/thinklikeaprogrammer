package main

import "fmt"

func main() {
	n := 5
	a := 1

	for i := n - 1; i >= 0 && i <= n-1; i -= a {
		if i == 0 {
			a = -1
		}
		for j := n - i; j >= 1; j-- {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

// a := 1

// for i := -n + 1; i <= 0 && i >= -n+1; i += a {
// 	// fmt.Println(i, n)
// 	if i == 0 {
// 		a = -1
// 	}
// 	for j := n + i; j >= 1; j-- {
// 		// fmt.Print(j)
// 		fmt.Print("#")
// 	}
// 	fmt.Println()
// }
