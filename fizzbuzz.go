// Write a short program that prints each number from 1 to 100 on a new line. 
// For each multiple of 3, print "Fizz" instead of the number. 
// For each multiple of 5, print "Buzz" instead of the number. 
// For numbers which are multiples of both 3 and 5, print "FizzBuzz" instead of the number.


package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0{
			fmt.Printf("FizzBUzz")
		}else if i%3 == 0 {
			fmt.Printf("Fizz")
		}else if i%5 == 0 {
			fmt.Printf("Buzz")
		}else{
			fmt.Printf("%d",i)
		}
		// A trailing new line (so both fizz + buzz can be printed on the same line)
		fmt.Printf("\n")
	}
}
