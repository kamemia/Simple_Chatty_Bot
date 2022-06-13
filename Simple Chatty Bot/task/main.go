package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello! My name is Aid.")
	fmt.Println("I was created in 2020.")
	fmt.Println("Please, remind me your name.")

	var name string
	var age, remainder3, remainder5, remainder7 int
	fmt.Scan(&name)

	fmt.Println("What a great name you have, " + name + "!")
	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")
	fmt.Scan(&remainder3, &remainder5, &remainder7)
	age = (remainder3*70 + remainder5*21 + remainder7*15) % 105
	// reading all remainders

	fmt.Println("Your age is " + strconv.Itoa(age) + "; that's a good time to start programming!")

	var num int

	fmt.Println("Now I will prove to you that I can count to any number you want.")
	fmt.Scan(&num)
	for i := 0; i <= num; i++ {
		fmt.Printf("%d!\n", i)
	}
	var n int
	fmt.Println("Let's test your programming knowledge.")
	fmt.Println("Why do we use methods?")
	fmt.Println("1. To repeat a statement multiple times.")
	fmt.Println("2. To decompose a program into several small subroutines.")
	fmt.Println("3. To determine the execution time of a program.")
	fmt.Println("4. To interrupt the execution of a program.")
	for {
		fmt.Scan(&n)
		if n == 2 {
			fmt.Println("Completed, have a nice day!")
			break
		} else {
			fmt.Println("Please, try again.")
		}
	}
	fmt.Println("Congratulations, have a nice day!")
}
