//Consider the following array declaration: nums := [...]int{30, -1, -6, 90, -6}

//Write a Go program that counts the number of positive even numbers in the array.

package main

import "fmt"

func main(){
	nums := [...]int{30, -1, -6, 90, -6}

	count := 0

	for _,v := range nums{
		if v >= 0 && v % 2 == 0{
			count++
		}
	}
	fmt.Println("number of positive even int :",count)
}
