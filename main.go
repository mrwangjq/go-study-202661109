package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	info, ok := debug.ReadBuildInfo()
	if ok {
		fmt.Println("编译使用的Go版本：", info.GoVersion)
	}
	fmt.Println("当前运行环境runtime版本：", runtime.Version())
	fmt.Println("Hello World")
}
