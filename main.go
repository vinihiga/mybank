package main

import "fmt"

func removeDuplicates(x *[]int) {
	dict := make(map[int]int)
	var i = 0

	for i < len(*x) {
		_, ok := dict[(*x)[i]]

		if ok {
			slice := make([]int, len(*x)-1)
			copy(slice, (*x)[:i])
			copy(slice[i:], (*x)[i+1:])
			*x = slice
			i--
		} else {
			dict[(*x)[i]] = 1
		}

		i++
	}
}

func main() {
	x := make([]int, 0, 4)
	x = append(x, 1, 1, 2, 3)
	removeDuplicates(&x)
	fmt.Println(x)
}
