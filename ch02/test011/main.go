package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var longestLength int
	var longestVowelsLength int
	var input string

	io := bufio.NewReader(os.Stdin)

	fmt.Println("Enter string")
	input, err := io.ReadString('\n')
	if err != nil {
		fmt.Println("Error in reading the string", err)
	}
	// fmt.Scanln(&input)
	fmt.Println("input", input)
	input = strings.Trim(input, "\n")
	inputArr := strings.Split(input, " ")
	for _, v := range inputArr {
		l, lv := getStats(v)
		if l > longestLength {
			longestLength = l
		}
		if lv > longestVowelsLength {
			longestVowelsLength = lv
		}
	}
	fmt.Println("longestLength", longestLength)
	fmt.Println("longestVowelsLength", longestVowelsLength)
}

func getStats(input string) (length, numberOfVowels int) {
	for _, v := range input {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			numberOfVowels++
		}
	}
	return len(input), numberOfVowels
}
