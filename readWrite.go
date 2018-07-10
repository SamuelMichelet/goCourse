package main

import (
	"io"
	"log"
	"os"
)

func main() {
	original, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer original.Close()

	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesWritten, err := io.Copy(newFile, original)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("copied %d bytes.", bytesWritten)

	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
