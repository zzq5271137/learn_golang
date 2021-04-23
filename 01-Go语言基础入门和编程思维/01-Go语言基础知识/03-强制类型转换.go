package main

import (
	"fmt"
	"math"
)

/**
 * 强制类型转换
 *
 * 注: golang没有隐式类型转换, 只有强制类型转换, 即所有的类型转换都必须显式地写出来
 */

func triangle() {
	var a, b int = 3, 4
	var c int
	//c = math.Sqrt(a*a + b*b) // math.Sqrt() 函数需要参数类型的是 float64, 所以这里报错 (因为golang中的类型转换是强制的)
	//c = math.Sqrt(float64(a*a + b*b)) // 这里如果直接赋值, 则还是报错, 因为 c 声明的是 int 类型 (因为golang中的类型转换是强制的)
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	triangle()
}
