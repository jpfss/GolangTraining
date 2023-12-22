package main

import "fmt"

func main() {
	a := "stored in a"
	//b申明并未使用也会报错
	b := "stored in b"
	fmt.Println("a - ", a)
	// b is not being used - invalid code
}
