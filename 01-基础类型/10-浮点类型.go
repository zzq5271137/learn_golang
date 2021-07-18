/**
 * 浮点类型
 */

package main

import "fmt"

func floatTest() {
	fmt.Println("floatTest():")

	/*
	 * float32: 单精度浮点型, 4字节, 小数位精确到7位
	 * float64: 双精度浮点型, 8字节, 小数位精确到15位
	 */
	var a float32 = 123.555555
	var b float64 = 123.555555
	fmt.Println(a)
	fmt.Println(b)

	c := 3.14
	fmt.Printf("%T\n", c) // 通过自动推导类型声明的浮点型变量, 默认为float64
}

func main() {
	floatTest()
}
