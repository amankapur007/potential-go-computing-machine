package main

import "fmt"

func main() {
	// declare
	var a = "initial"
	fmt.Println(a)

	//number
	var b, c int = 1, 2
	fmt.Println(b, c)

	//boolean
	var d = true
	fmt.Println(d)

	//defaultusing this we can decalare with out initialization
	var e int
	fmt.Println(e)

	// if we using short hand then initialization is required
	f := "apple"
	fmt.Println(f)
}
