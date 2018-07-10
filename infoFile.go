package main

import(
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	 }

	log.Println("file does exist;File information: ")
	log.Println(fileInfo)

	
	fmt.Println("\n\nfile name:", fileInfo.Name())
	fmt.Println("size bytes::", fileInfo.Size())
	fmt.Println("file permission:", fileInfo.Mode())
	fmt.Println("last modified:", fileInfo.ModTime())
	fmt.Println("is directory:", fileInfo.IsDir())
	fmt.Printf("systeme interface type: %T\n ", fileInfo.Sys())
	fmt.Printf("system info: %+v\n\n ", fileInfo.Sys())

}




