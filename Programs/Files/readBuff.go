package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file,err := os.OpenFile("test.txt",os.O_CREATE,0644)
	if err != nil{
		log.Fatal(err)
	}
    defer file.Close()

	byteSlice := make([]byte,2) // we are reading 2 bytes in the file
	numberBytesRead,err := io.ReadFull(file,byteSlice)
	if err != nil{
		log.Fatal(err)
	}
	log.Printf("Number bytes read:%d\n",numberBytesRead)
	log.Printf("data read:%s\n",byteSlice)

	// now we see how to read full file
	file,err = os.Open("main.go")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	data,err := ioutil.ReadAll(file)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("Data as string:%s\n",data)
	fmt.Printf("number of lines read:%d\n",len(data))

	data,err = ioutil.ReadFile("test.txt")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("data read:%s\n",data)
}

