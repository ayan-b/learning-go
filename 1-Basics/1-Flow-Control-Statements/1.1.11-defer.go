/*

Defer

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
*/

package main

import "fmt"

func main() {
	fmt.Println("hello1")
	defer fmt.Println("world1")

	fmt.Println("hello2")
	defer fmt.Println("world2")
	defer fmt.Println("world3")
	/*
	hello1
	hello2
	world3
	world2
	world1
	*/
}
