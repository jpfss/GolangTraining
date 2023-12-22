package main

import (
	"fmt"
	//这里导入的时候应该采用module名称而不是目录名
	"pack/stringutil"
	"pack/winniepooh"
	//"github.com/jpfss/GolangTraining/02_package/stringutil"
	//someAlias "github.com/jpfss/GolangTraining/02_package/icomefromalaska"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
	fmt.Println(winniepooh.BearName)
}
