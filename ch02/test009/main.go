package main

import "fmt"

func main() {

	input := []byte{}
	var c byte
	fmt.Scanf("%c", &c)
	for c == '0' || c == '1' {
		input = append(input, c)
		fmt.Scanf("%c", &c)
	}

	fmt.Println("decimal = ", binaryToDecimal(input))
	hex, err := binaryToHex(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("hexadecimal = ", hex)
}

func binaryToDecimal(input []byte) int {
	var output int

	for _, v := range input {
		output *= 2
		output += int(v) - '0'
	}
	return output
}

func binaryToHex(input []byte) ([]int, error) {
	output := []int{}

	l := len(input)
	if l%3 != 0 {
		return nil, fmt.Errorf("input is not a hex")
	}

	for i := 0; i <= l-3; i += 3 {
		d := binaryToDecimal([]byte{input[i], input[i+1], input[i+2]})
		output = append(output, d)
	}
	return output, nil
}
