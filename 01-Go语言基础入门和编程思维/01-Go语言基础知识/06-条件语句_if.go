package main

import (
	"fmt"
	"io/ioutil"
)

/**
 * 条件语句: if
 */

func bounded(v int) int {
	if v > 100 {
		return 100
	} else if v < 0 {
		return 0
	} else {
		return v
	}
}

func main() {
	fmt.Println(bounded(101))

	//const fileName = "06-条件语句_if.txt"  // 由于.go源文件编译后与.go源文件可能不处在同一目录下, 所以如果是这种方式，可能找不到文件
	const fileName = "/Users/zzq/Documents/workspaces/GoLand_workspace/learn_golang/01-Go语言基础入门和编程思维/01-Go语言基础知识/06-条件语句_if.txt"
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents) // contents是一个byte数组, 使用 %s 可以打出来
	}

	// if 的条件里也可以进行赋值, 赋值完后再进行判断
	// if 的条件里赋值的变量的作用域就在这个 if 语句里
	if contents, err := ioutil.ReadFile(fileName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
