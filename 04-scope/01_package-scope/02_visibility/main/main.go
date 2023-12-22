package main

import (
	"fmt"
	//改为本地包引用，module模式下从go.mod文件目录开始，但是要以module名而不是目录名开头
	"scope/01_package-scope/02_visibility/vis"
	//"github.com/jpfss/GolangTraining/04_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
