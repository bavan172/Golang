package main

import (
	"bufio"
	"log"
	"os"
)

func main (){
	file,err:=os.OpenFile("my_file.txt",os.O_WRONLY|os.O_CREATE,0644)
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	// create buffer writer using bufio
	buffWritter := bufio.NewWriter(file)
	bs := []byte{97,99,98}
	byteWritten,err := buffWritter.Write(bs)
	if err != nil{
		log.Fatal(err)
	}
	log.Printf("bytes written to buffer (not file):%d\n",byteWritten)

	bytesAvailable := buffWritter.Available()
	log.Printf("bytes available in buffer:%d\n",bytesAvailable)

	// writing string
	byteWritten,err = buffWritter.WriteString("\nJust random string")
	if err != nil{
		log.Fatal(err)
	}
	// data stored in buffer waiting to be written into disk
	unflushedBufSize := buffWritter.Buffered()
	log.Printf("bytes flushed:%d\n",unflushedBufSize)

	//finally to write data into the file from buffer
	buffWritter.Flush()

}
