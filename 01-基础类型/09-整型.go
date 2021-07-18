/**
 * 整型
 */

package main

import (
	"fmt"
	"unsafe"
)

func intTest() {
	fmt.Println("intTest():")

	/*
	 * int: 有符号整型
	 * uint: 无符号整型
	 */
	var a int = -1
	//var b uint = -10 // 运行报错
	var b uint = 1
	fmt.Println(a, b)
	fmt.Printf("type of a: %T, size of a: %d\n", a, unsafe.Sizeof(a)) // 根据操作系统属于32位还是64位, int所占字节数为4或8

	// 或者, 在声明变量时可以指定所占位数
	// 当然, 每种类型都有其取值范围, 例如, int8: -128~127, unit8: 0~255
	var c int8 = 1
	var d int16 = 1
	var e int32 = 1
	var f int64 = 1
	fmt.Printf("type of c: %T, size of c: %d\n", c, unsafe.Sizeof(c))
	fmt.Printf("type of d: %T, size of d: %d\n", d, unsafe.Sizeof(d))
	fmt.Printf("type of e: %T, size of e: %d\n", e, unsafe.Sizeof(e))
	fmt.Printf("type of f: %T, size of f: %d\n", f, unsafe.Sizeof(f))
}

func main() {
	intTest()
}
