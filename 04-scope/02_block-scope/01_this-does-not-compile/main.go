package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo()
}

func foo() {
	// no access to x
	// this does not compile
	//这里会报错-x没有定义
	fmt.Println(x)
}
