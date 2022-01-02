package main

import (
	"fmt"
	"sort"
)

func main() {
	var myIntArr MyIntArr = []int{10, 5, 1}
	sort.Sort(myIntArr)
	fmt.Println(myIntArr)
}

type MyIntArr []int

func (myIntArr MyIntArr) Len() int {
	return len(myIntArr)
}

func (myIntArr MyIntArr) Less(i, j int) bool {
	return int(myIntArr[i]) < int(myIntArr[j])
}

func (myIntArr MyIntArr) Swap(i, j int) {
	myIntArr[i], myIntArr[j] = myIntArr[j], myIntArr[i]
}
