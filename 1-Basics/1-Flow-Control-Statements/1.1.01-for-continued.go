/*
For continued

The init and post statements are optional.
*/

package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
	// for ; ; {
	// 	sum += sum
	// 	break
	// }
	// fmt.Println(sum)
}
