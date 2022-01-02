package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	var b int = (n / 2) - 1
	fmt.Println()
	var direction int

	for i := 2; i <= n && i >= 2; {

		printSpace(b)
		printHash(i)
		printSpace(b)

		if direction == 0 {
			i += 2
			b--
		} else {
			i -= 2
			b++
		}

		if i == n {
			direction = 1
		}

		fmt.Println()
	}
}

func printHash(n int) {
	for j := 0; j < n; j++ {
		fmt.Print("#")
	}
}

func printSpace(b int) {
	for j := 0; j < b; j++ {
		fmt.Print(" ")
	}
}
