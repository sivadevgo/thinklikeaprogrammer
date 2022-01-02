package main

import "fmt"

func main() {

	var c byte
	var n int
	var sum int
	var lastChar byte
	fmt.Scanf("%c\n", &c)
	fmt.Println(c)
	for c >= 48 && c < 58 {
		lastChar = c
		if n%2 == 1 {
			sum += doubleTheValue(int(c) - '0')
		} else {
			sum += int(c) - '0'
		}
		fmt.Println("sum", sum)
		n++
		fmt.Scanf("%c\n", &c)
	}

	if n%2 == 0 {
		sum -= doubleTheValue(int(lastChar) - '0')
	} else {
		sum -= (int(lastChar) - '0')
	}
	fmt.Println("sum", sum)
	fmt.Println("lastChar", lastChar-'0')
	sum += (int(lastChar) - '0')
	fmt.Println("sum", sum)
	if sum%10 == 0 {
		fmt.Println("SUCCESS")
	} else {
		fmt.Println("FAILED")
	}

}

func doubleTheValue(input int) int {
	d := input * 2
	if d < 10 {
		return d
	} else {
		return d%10 + 1
	}
}
