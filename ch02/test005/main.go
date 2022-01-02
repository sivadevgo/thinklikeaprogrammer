package main

import (
	"fmt"
)

func main() {

	var c byte

	fmt.Scanf("%c", &c)
	var currentNumber int
	var currentState int //0 - Upper, 1 - Lower, 2 - Punctuation
	var output []byte

	for c >= 48 && c < 58 || c == ',' || c == '\n' {
		fmt.Printf("%c\n", c)
		switch c {
		case ',', '\n':
			fmt.Println("currentNumber : ", currentNumber)
			var mod int
			if currentState == 2 {
				mod = currentNumber % 9
			} else {
				mod = currentNumber % 27
			}
			if mod == 0 {
				fmt.Println("mod==0")
				currentState = changeState(currentState)
				fmt.Println("changed state to ", currentState)
			} else if currentState == 0 {
				fmt.Println("state==0")
				l := getUpperCaseLetter(mod)
				fmt.Printf("%c\n", l)
				output = append(output, l)
			} else if currentState == 1 {
				fmt.Println("state==1")
				l := getLowerCaseLetter(mod)
				fmt.Printf("%c\n", l)
				output = append(output, l)
			} else if currentState == 2 {
				fmt.Println("state==2")
				l := getPunctuation(mod)
				fmt.Printf("%c\n", l)
				output = append(output, l)
			}

			currentNumber = 0
		default:
			currentNumber = (currentNumber * 10) + (int(c) - '0')
		}
		if c == '\n' {
			fmt.Println("******")
			break
		}
		fmt.Scanf("%c", &c)
	}
	fmt.Println("*", c)
	print(output)
	fmt.Printf("*%c*", '\n')
}

func print(input []byte) {
	for _, c := range input {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

func getUpperCaseLetter(input int) byte {
	return 'A' + byte(input) - 1
}
func getLowerCaseLetter(input int) byte {
	return 'a' + byte(input) - 1
}

func changeState(input int) int {
	return (input + 1) % 3
}

func getPunctuation(input int) byte {
	switch input {
	case 1:
		return '!'
	case 2:
		return '?'
	case 3:
		return ','
	case 4:
		return '.'
	case 5:
		return ' '
	case 6:
		return ';'
	case 7:
		return '"'
	case 8:
		return '\''
	}
	return 0
}
