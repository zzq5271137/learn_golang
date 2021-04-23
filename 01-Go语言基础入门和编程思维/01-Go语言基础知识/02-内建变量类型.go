package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
 * 内建变量类型:
 * 1. 布尔型: bool
 * 2. 字符型: rune  (不同于Java中的Char是一个字节大小, rune是4个字节大小, 即32位)
 * 3. 字符串型: string
 * 4. 整数型:
 *    (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
 *    注: 前面加 (u) 表示无符号整数, 不加表示有符号整数
 * 5. 字节型: byte  (8位)
 * 6. 浮点数类型:
 *    float32, float64, complex64, complex128
 *    注: complex表示复数
 */

func complexNumber() {
	c := 3 + 4i // go语言的复数的表示
	fmt.Println(cmplx.Abs(c))
}

/**
 * @Description: 模拟欧拉公式
 */
func euler() {
	/*
	 * 并不正好是0, 因为go语言中的复数:
	 * 1. complex64, 它的实部和虚部都是一个float32的数
	 * 2. complex128, 它的实部和虚部都是一个float64的数
	 * 即, 它的实部和虚部的浮点数并不是完全准的
	 */
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1) // 这种写法更好, 因为 cmplx.Exp() 就是表示 e 的多少次方

	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1) // 只要小数点后三位, 这样结果是0
}

func main() {
	complexNumber()
	euler()
}
