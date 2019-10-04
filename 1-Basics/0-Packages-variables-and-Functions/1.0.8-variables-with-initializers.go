/*
Variables with initializers

A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
*/

package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = false, false, "no!"
	fmt.Println(c, python, java, i, j)
}
