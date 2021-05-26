// 1. Create an anonymous function that calculates and prints out the square root of a float value it receives as argument.
// 2. Launch the function as a goroutine and synchronize it with main using WaitGroups

package main

import (
	"fmt"
	"math"
	"sync"
)

func sum(a,b float64, wg *sync.WaitGroup){
	s := a + b
	fmt.Printf("sum of %.2f and %.2f is %.2f\n",a,b,s)
	wg.Done()
}

func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go func(n float64, wg *sync.WaitGroup){
		x := math.Sqrt(n)
		fmt.Printf("Square root of %.2f is %.4f\n",n,x)
		wg.Done()
	}(16.1,&wg)
	wg.Wait()
}
