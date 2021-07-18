/**
 * bool类型
 */

package main

import "fmt"

func boolTest() {
	fmt.Println("boolTest():")

	var v1, v2 bool // bool类型的Zero Value为false
	v1 = true
	v3 := (1 == 2)
	fmt.Println(v1, v2, v3)

	// bool类型不能接收其他类型的赋值, 也不支持自动或强制的类型转换
	//var b bool
	//b = 1       // 编译错误
	//b = bool(1) // 编译错误
}

func main() {
	boolTest()
}
