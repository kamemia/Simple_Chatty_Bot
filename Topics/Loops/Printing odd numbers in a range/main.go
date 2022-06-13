package main

import "fmt"

func main() {
	// put your code her
	var i = 2

	for ; i < 5; i++ {
		if i == 4 {
			break
		}
	}

	fmt.Print(i)
}
