package main

import "fmt"

func main() {
	// DO NOT delete or modify the code block below!
	var s1 = []int{1, 2, 3, 4, 5, 6}
	var s2 []int

	var n = copy(s2, s1)
	fmt.Println(n)
}
