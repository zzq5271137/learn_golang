/**
 * 常量
 */

package main

import "fmt"

func defineConst() {
	fmt.Println("defineConst():")

	/*
	 * 常量使用 const 关键字定义
	 */
	const a int = 2
	const ( // 同时定义多个常量
		b float64 = 3.2
		c         = "hello" // 自动推导类型, 省略类型声明
	)
	fmt.Println(a, b, c)
}

func main() {
	defineConst()
}
