package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("hy!!")
	 s, sep := "", ""
	for a := range os.Args[1:] {
		s += sep + a
		sep = " "
	}
	fmt.Println(os.Args[a])
}
