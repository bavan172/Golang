// writing using  os and ioutil

package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main(){
	file,err := os.OpenFile("a.txt",os.O_WRONLY|os.O_CREATE,0644)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("I learn Golang")
	// writing byte slice to file
	byteWritten,err := file.Write(byteSlice)
	if err != nil{
		log.Fatal(err)
	}
	
  // displaying number of bytes written to the file
	log.Printf("Bytes written:%d",byteWritten)

	// creating another byte slice
	bs := []byte("Go is cool")
	err = ioutil.WriteFile("c.txt",bs,0644)
	if err != nil{
		log.Fatal(err)
	}
}
	
