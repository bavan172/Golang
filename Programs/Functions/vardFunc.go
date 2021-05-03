//Create a function with the identifier sum 
// takes in a variadic parameter of type int 
// returns the sum of all values of type int passed in.

package main

import "fmt"

func sum(a ...int)int{
	s := 0
	for _,v := range a{
		s += v
	}
	return s
}

func main() {
	n := sum(1,2,4,5,2,3,5)
	fmt.Println(n)
}
