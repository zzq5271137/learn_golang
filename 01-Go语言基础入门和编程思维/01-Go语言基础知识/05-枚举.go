package main

import "fmt"

/**
 * Golang的枚举
 */

func enums() {
	//const ( // go语言没有单独的枚举类型关键字, 使用 const 加括号的形式定义枚举类型
	//	cpp    = 0
	//	java   = 1
	//	python = 2
	//	golang = 3
	//)
	const (
		cpp = iota // 可以使用这种方式来定义, 表示枚举的值为自增的 (iota 是一个自增值的种子, 从0开始)
		java
		_ // 表示占位
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)

	// b, kb, mb, gb, tb
	const (
		b = 1 << (10 * iota) // <<: 位运算, 左移
		kb
		mb
		gb
		tb
	)
	fmt.Println(b, kb, mb, gb, tb)
}

func main() {
	enums()
}
