package main

import "fmt"

func main() {
	var c byte
	var number int

	fmt.Scanf("%c", &c)

	for c >= 48 && c < 58 {
		number = (number * 10) + (int(c) - 48)
		fmt.Scanf("%c", &c)
	}
	fmt.Println(number)
}
