// Program to find factorial of a number

package main

import "fmt"

func fact(n int , c chan int){
	f := 1
	for i:=2;i<=n;i++{
		f *= i
	}
	c <- f
}

func main(){
	ch := make(chan int)
	defer close(ch)

	go fact(5,ch)

	f := <- ch
	fmt.Println(f)
	}

