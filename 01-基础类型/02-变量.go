/**
 * 变量
 */

package main

import "fmt"

func defineVariable() {
	fmt.Println("defineVariable1():")

	/*
	 * 变量的声明:
	 * 1. 声明方式:
	 *    var variableName variableType
	 * 2. 在golang中, 变量声明后就必须被使用
	 * 3. 在golang中, 只是声明变量但是没有赋初始值时, 会有一个Zero Value, 例如, int的Zero Value是0, string的Zero Value是空字符串
	 */
	var num int
	var str string
	fmt.Printf("num=%d, str=%q\n", num, str) // %q: 打印字符串, 并附带双引号

	var a, b float64 = 2.1, 3.2 // 可以同时声明多个变量, 并赋初始值
	fmt.Println(a, b)

	var c = 2 // 如果赋了初始值, 可以省略类型声明, 编译器会自动推导出类型
	fmt.Println(c)

	var ( // 也可以这么定义多个变量
		d bool       // 只定义不赋值 (bool的Zero Value是false)
		e int  = 2   // 定义并赋值
		f      = 2.7 // 自动推导类型, 省略类型声明
	)
	fmt.Println(d, e, f)
}

func main() {
	defineVariable()
}
