// Print Hello, World. on the first line, and the contents of input string on the second line

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin) 
    
    scanner.Scan()
    
    text := scanner.Text()
    fmt.Println("Hello, World.")
    fmt.Println(text)
}
