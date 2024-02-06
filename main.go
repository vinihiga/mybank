package main

import "fmt"

func removeDuplicates(x *[]int) {
	dict := make(map[int]int)
	var i = 0

	for i < len(*x) {
		_, ok := dict[(*x)[i]]

		if ok {
			*x = append((*x)[:i], (*x)[i+1:]...)
			i--
		} else {
			dict[(*x)[i]] = 1
		}

		i++
	}
}

func main() {
	x := make([]int, 0, 4)
	x = append(x, 0, 0, 1, 0, 0, 1)
	removeDuplicates(&x)
	fmt.Println(x)
}
