package main

import "fmt"

func main() {
	var a = [4][4]int{
		{13, 56, 1, -2},
		{-34, 1, 70, -23},
		{6, 11, -1, 0},
		{-67, 1, 9, 2},
	}
	fmt.Println(a[2][0])
}
