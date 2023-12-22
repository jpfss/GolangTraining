package main

import "fmt"

func main() {
	//未定义先使用变量无法正常执行
	fmt.Println(x)
	fmt.Println(y)
	x := 42
}

var y = 42
