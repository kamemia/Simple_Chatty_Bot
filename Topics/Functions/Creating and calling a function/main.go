package main

import "fmt"

func name(name string) string {
	// write the function code block here
	return name + " is learning how to call functions!"
}

func main() {
	// Do not change this two lines of code
	var userName string
	fmt.Scanf("%s", &userName)

	// call the function directly, or within a fmt.Println statement here.
	fmt.Println(name(userName))

}
